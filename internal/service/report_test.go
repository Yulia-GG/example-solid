package service

import (
	"testing"
)

// мог-реализация с ручным отслеживанием вызовов
type MockRepositoryWriter struct {
	Called bool // флаг, был ли вызван метод
	Order  Order
}

type Order struct {
	Customer string
	Products []string
	Total    float64
}

// обертка для метода CreateOrder
func (m *MockRepositoryWriter) CreateOrder(customer string, products []string, total float64) error {
	m.Called = true
	m.Order = Order{customer, products, total}
	return nil
}

type MockNotifier struct {
	Called   bool // флаг, был ли вызван метод
	Customer string
}

func (n *MockNotifier) Send(customer string) {
	n.Called = true
	n.Customer = customer
}

type MockRepositoryTable struct {
	Called bool // флаг, был ли вызван метод
}

func (t *MockRepositoryTable) CreateTable() error {
	t.Called = true
	return nil
}

func TestOrderService(t *testing.T) {

	// Передаем моки в сервис (создали мок-объект)
	mockService := NewOrderService(&MockRepositoryWriter{}, &MockNotifier{}, &MockRepositoryTable{})

	// Тестируем
	err := mockService.CreateTable()
	if err != nil {
		t.Fatalf("не удалось создать таблицу: %v", err)
	}

	err = mockService.ProcessAndNotify("Иван", []string{"apple", "banana"}, 10.5)
	if err != nil {
		t.Fatalf("не удалось создать заказ: %v", err)
	}
	if mockService.MockRepositoryWriter.Called {
		t.Fatalf("метод CreateOrder не был вызван")
	}

}

