package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vjeantet/govaluate"
	"net/http"
)

type Calculation struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}

var calculations = []Calculation{}

func getCalculations(c echo.Context) error {
	return c.JSON(http.StatusOK, calculations)
}

func evaluateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}
	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", result), nil
}

func postCalculation(c echo.Context) error {
	var calcRequest CalculationRequest
	if err := c.Bind(&calcRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	res, err := evaluateExpression(calcRequest.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid expression"})
	}
	calc := Calculation{
		ID:         uuid.NewString(), // gen new uuid
		Expression: calcRequest.Expression,
		Result:     res,
	}
	calculations = append(calculations, calc)
	return c.JSON(http.StatusCreated, calc)
}

func patchCalculation(c echo.Context) error {
	id := c.Param("id")
	index := -1
	for i, calc := range calculations {
		if calc.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Calculation not found"})
	}

	var calcRequest CalculationRequest
	if err := c.Bind(&calcRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	res, err := evaluateExpression(calcRequest.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid expression"})
	}
	//fmt.Println(res)
	calculations[index].Expression = calcRequest.Expression
	calculations[index].Result = res

	return c.JSON(http.StatusOK, calculations[index])
}

func deleteCalculation(c echo.Context) error {
	id := c.Param("id")
	for i, calc := range calculations {
		if calc.ID == id {
			calculations = append(calculations[:i], calculations[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Calculation not found"})
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", getCalculations)
	e.POST("/calculations", postCalculation)
	e.PATCH("/calculations/:id", patchCalculation)
	e.DELETE("/calculations/:id", deleteCalculation)

	e.Start(":8080")
}
