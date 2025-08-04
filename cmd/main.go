package main

import (
	"CalculatorAppBackend/internal/db"
	"CalculatorAppBackend/internal/handlers"
	"CalculatorAppBackend/internal/services/calculationService"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	e := echo.New()

	calcRepository := calculationService.NewCalculationRepository(database)
	calcService := calculationService.NewCalculationService(calcRepository)
	calcHandlers := handlers.NewCalculationHandler(calcService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", calcHandlers.GetCalculations)
	e.POST("/calculations", calcHandlers.PostCalculation)
	e.PATCH("/calculations/:id", calcHandlers.PatchCalculation)
	e.DELETE("/calculations/:id", calcHandlers.DeleteCalculation)

	if err = e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
