package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tasukuwatanabe/crm-app/internal/handler"
	"github.com/tasukuwatanabe/crm-app/internal/repository"
	"github.com/tasukuwatanabe/crm-app/internal/usecase"
)

func main() {
	e := echo.New()

	customerRepo := repository.NewInMemoryCustomerRepository()
	customerUC := usecase.NewCustomerUsecase(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerUC)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	g := e.Group("/api/v1")
	g.GET("/customers", customerHandler.GetAll)
	g.GET("/customers/:id", customerHandler.GetByID)
	g.POST("/customers", customerHandler.Create)
	g.PUT("/customers/:id", customerHandler.Update)
	g.DELETE("/customers/:id", customerHandler.Delete)

	e.Logger.Fatal(e.Start(":8080"))
}
