package orders

import (
	"order-management-system/internal/orders/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	orderGroup := app.Group("/orders")

	orderGroup.Get("/", handlers.GetOrders)  // Obtener todos los pedidos
	orderGroup.Get("/:id", handlers.GetOrderByID)  // Obtener un pedido por ID
	orderGroup.Post("/", handlers.CreateOrder)  // Crear un pedido
	orderGroup.Put("/:id", handlers.UpdateOrder)  // Actualizar un pedido completo
	orderGroup.Patch("/:id/status", handlers.UpdateOrderStatus)  // Actualizar estado de un pedido
	orderGroup.Delete("/:id", handlers.DeleteOrder)  // Eliminar un pedido
}
