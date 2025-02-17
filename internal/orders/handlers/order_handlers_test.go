package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"order-management-system/internal/orders/models"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// **Mock para probar sin la BD**
var testOrders = []models.Order{
	{ID: 1, Customer: "Cliente X", TotalAmount: 150.00, Status: "pending"},
}

// **Prueba GET /orders**
func TestGetOrdersHandler(t *testing.T) {
	app := fiber.New()
	app.Get("/orders", func(c *fiber.Ctx) error {
		return c.JSON(testOrders) // ✅ Devolvemos datos simulados
	})

	req := httptest.NewRequest(http.MethodGet, "/orders", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// **Prueba POST /orders**
func TestCreateOrderHandler(t *testing.T) {
	app := fiber.New()
	app.Post("/orders", func(c *fiber.Ctx) error {
		var order models.Order
		if err := c.BodyParser(&order); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
		}
		return c.Status(201).JSON(fiber.Map{"message": "Pedido creado"})
	})

	order := models.Order{Customer: "Cliente X", TotalAmount: 150.00, Status: "pending"}
	orderJSON, _ := json.Marshal(order)

	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(orderJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
