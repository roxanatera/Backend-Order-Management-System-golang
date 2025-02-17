package handlers

import (
	"order-management-system/internal/orders/models"
	"order-management-system/internal/orders/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetOrders(c *fiber.Ctx) error {
	orders, err := services.GetAllOrders()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error obteniendo pedidos"})
	}
	return c.JSON(orders)
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}
	err := services.CreateNewOrder(order)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al crear pedido"})
	}
	return c.Status(201).JSON(fiber.Map{"message": "Pedido creado"})
}


// Obtener un pedido por ID
func GetOrderByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}

	order, err := services.GetOrder(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Pedido no encontrado"})
	}

	return c.JSON(order)
}

// Actualizar un pedido completo (PUT)
func UpdateOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}

	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	order.ID = id
	err = services.UpdateExistingOrder(order)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error actualizando pedido"})
	}

	return c.JSON(fiber.Map{"message": "Pedido actualizado"})
}

// Actualizar solo el estado del pedido (PATCH)
func UpdateOrderStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}

	var request struct {
		Status string `json:"status"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	err = services.UpdateOrderStatus(id, request.Status)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error actualizando estado"})
	}

	return c.JSON(fiber.Map{"message": "Estado actualizado"})
}

// Eliminar un pedido (DELETE)
func DeleteOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}

	err = services.DeleteOrder(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error eliminando pedido"})
	}

	return c.JSON(fiber.Map{"message": "Pedido eliminado"})
}

