package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tasukuwatanabe/crm-app/internal/db"
	"github.com/tasukuwatanabe/crm-app/internal/handler"
	"github.com/tasukuwatanabe/crm-app/internal/repository"
	"github.com/tasukuwatanabe/crm-app/internal/usecase"
)

func main() {
	conn, err := db.NewConnection(db.DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "tasuku",
		Password: "",
		DBName:   "crm_development",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer conn.Close()

	e := echo.New()

	customerRepo := repository.NewPostgresCustomerRepository(conn)
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
