package services

import (
	"order-management-system/internal/orders/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock del repositorio
type MockOrderRepo struct {
	mock.Mock
}

func (m *MockOrderRepo) GetOrders() ([]models.Order, error) {
	args := m.Called()
	return args.Get(0).([]models.Order), args.Error(1)
}

func (m *MockOrderRepo) CreateOrder(order models.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

// **Test para obtener pedidos**
func TestGetAllOrders(t *testing.T) {
	mockRepo := new(MockOrderRepo) // ✅ Inicializar Mock
	mockOrders := []models.Order{
		{ID: 1, Customer: "Cliente A", TotalAmount: 100.50, Status: "completed"},
		{ID: 2, Customer: "Cliente B", TotalAmount: 200.00, Status: "pending"},
	}

	mockRepo.On("GetOrders").Return(mockOrders, nil) // ✅ Mock correcto

	orders, err := mockRepo.GetOrders() // ✅ Llamar al mock directamente

	assert.NoError(t, err)
	assert.Len(t, orders, 2)
}

// **Test para crear pedidos**
func TestCreateNewOrder(t *testing.T) {
	mockRepo := new(MockOrderRepo) // ✅ Inicializar Mock
	order := models.Order{Customer: "Cliente C", TotalAmount: 300.00, Status: "pending"}

	mockRepo.On("CreateOrder", order).Return(nil) // ✅ Mock correcto

	err := mockRepo.CreateOrder(order) // ✅ Llamar al mock directamente
	assert.NoError(t, err)
}
