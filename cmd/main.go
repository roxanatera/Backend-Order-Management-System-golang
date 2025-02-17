package main

import (
	"log"
	"order-management-system/internal/orders"
	"order-management-system/pkg/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// Habilitar CORS para permitir peticiones desde el frontend
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Cambiar esto si el frontend tiene un dominio específico
		AllowMethods: "GET,POST,PUT,PATCH,DELETE",
		AllowHeaders: "Content-Type",
	}))

	// Ruta de prueba para ver si funciona la conexión
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("✅ API en ejecución")
	})

	// Conectar a la base de datos
	db.InitDB()

	// Configurar rutas de pedidos
	orders.SetupRoutes(app)

	// Iniciar servidor en el puerto 8080
	log.Fatal(app.Listen(":8080"))
}
