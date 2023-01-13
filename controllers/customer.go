package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sample-api-rest/models"
)

var CustomerController *customerController

type customerController struct {
}

type NewCustomerRequest struct {
	CustomerIdNumber    string `json:"customerIdNumber"`
	CustomerIdType      string `json:"customerIdType"`
	CustomerName        string `json:"customerName"`
	CustomerEmail       string `json:"customerEmail"`
	CustomerPhoneNumber string `json:"customerPhoneNumber"`
	CustomerAddress     string `json:"customerAddress"`
}

func (t *customerController) FindAll(c echo.Context) error {
	list, _ := models.Customer.FindAll()
	return c.JSON(http.StatusOK, list)
}

func (t *customerController) FindById(c echo.Context) error {
	list, _ := models.Customer.FindById(c.Param("id"))
	return c.JSON(http.StatusOK, list)
}
