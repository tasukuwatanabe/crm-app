package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tasukuwatanabe/crm-app/internal/domain"
	"github.com/tasukuwatanabe/crm-app/internal/usecase"
)

type CustomerHandler struct {
	uc *usecase.CustomerUsecase
}

func NewCustomerHandler(uc *usecase.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{uc: uc}
}

func (h *CustomerHandler) GetAll(c echo.Context) error {
	customers, err := h.uc.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, customers)
}

func (h *CustomerHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	customer, err := h.uc.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) Create(c echo.Context) error {
	var customer domain.Customer
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	created, err := h.uc.Create(&customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, created)
}

func (h *CustomerHandler) Update(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	var customer domain.Customer
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	customer.ID = id

	updated, err := h.uc.Update(&customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

func (h *CustomerHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	if err := h.uc.Delete(id); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
