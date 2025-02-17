package services

import (
	"order-management-system/internal/orders/models"
	"order-management-system/internal/orders/repository"
)


func GetAllOrders() ([]models.Order, error) {
	return repository.GetOrders()
}

func CreateNewOrder(order models.Order) error {
	return repository.CreateOrder(order)
}


// Obtener un pedido por ID
func GetOrder(id int) (models.Order, error) {
	return repository.GetOrderById(id)
}

// Actualizar un pedido (PUT)
func UpdateExistingOrder(order models.Order) error {
	return repository.UpdateOrder(order)
}

// Actualizar solo el estado del pedido (PATCH)
func UpdateOrderStatus(id int, status string) error {
	return repository.UpdateOrderStatus(id, status)
}

// Eliminar un pedido (DELETE)
func DeleteOrder(id int) error {
	return repository.DeleteOrder(id)
}
