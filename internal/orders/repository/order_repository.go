package repository

import (
	"order-management-system/internal/orders/models"
	"order-management-system/pkg/db"
)

func GetOrders() ([]models.Order, error) {
	rows, err := db.DB.Query("SELECT id, customer, total_amount, status FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.Customer, &order.TotalAmount, &order.Status); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func CreateOrder(order models.Order) error {
	_, err := db.DB.Exec("INSERT INTO orders (customer, total_amount, status) VALUES ($1, $2, $3)",
		order.Customer, order.TotalAmount, order.Status)
	return err
}

// Obtener un pedido por ID
func GetOrderById(id int) (models.Order, error) {
	var order models.Order
	err := db.DB.QueryRow("SELECT id, customer, total_amount, status FROM orders WHERE id = $1", id).
		Scan(&order.ID, &order.Customer, &order.TotalAmount, &order.Status)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

// Actualizar un pedido (PUT)
func UpdateOrder(order models.Order) error {
	_, err := db.DB.Exec("UPDATE orders SET customer=$1, total_amount=$2, status=$3 WHERE id=$4",
		order.Customer, order.TotalAmount, order.Status, order.ID)
	return err
}

// Actualizar parcialmente un pedido (PATCH)
func UpdateOrderStatus(id int, status string) error {
	_, err := db.DB.Exec("UPDATE orders SET status=$1 WHERE id=$2", status, id)
	return err
}

// Eliminar un pedido (DELETE)
func DeleteOrder(id int) error {
	_, err := db.DB.Exec("DELETE FROM orders WHERE id=$1", id)
	return err
}
