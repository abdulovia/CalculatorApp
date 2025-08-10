package handlers

import (
	"CalculatorAppBackend/internal/services/calculationService"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CalculationHandler struct {
	service calculationService.CalculationService
}

func NewCalculationHandler(service calculationService.CalculationService) *CalculationHandler {
	return &CalculationHandler{service: service}
}

func (h *CalculationHandler) GetCalculations(c echo.Context) error {
	calculations, err := h.service.GetAllCalculations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get calculations"})
	}
	return c.JSON(http.StatusOK, calculations)
}

func (h *CalculationHandler) PostCalculation(c echo.Context) error {
	var calcRequest calculationService.CalculationRequest
	if err := c.Bind(&calcRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	calc, err := h.service.CreateCalculation(calcRequest.Expression)
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not add calculation"})
	}

	return c.JSON(http.StatusCreated, calc)
}

func (h *CalculationHandler) PatchCalculation(c echo.Context) error {
	id := c.Param("id")

	var calcRequest calculationService.CalculationRequest
	if err := c.Bind(&calcRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	calc, err := h.service.UpdateCalculation(id, calcRequest.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not update calculation"})
	}

	return c.JSON(http.StatusOK, calc)
}

func (h *CalculationHandler) DeleteCalculation(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.DeleteCalculation(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not delete calculation"})
	}
	return c.NoContent(http.StatusNoContent)
}
