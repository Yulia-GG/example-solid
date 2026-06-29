package service

import (
	"testing"
	"reflect"
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

func (n *MockNotifier) Send(customer string) error {
	n.Called = true
	n.Customer = customer
	return nil
}

type MockRepositoryTable struct {
	Called bool // флаг, был ли вызван метод
}

func (t *MockRepositoryTable) CreateTable() error {
	t.Called = true
	return nil
}

func TestOrderService(t *testing.T) {

	// создали мок-объекты
	mockRepositoryWriter := &MockRepositoryWriter{}
	mockNotifier := &MockNotifier{}
	mockRepositoryTable := &MockRepositoryTable{}

	// Вызываем сервис и передаем моки
	mockService := NewOrderService(mockRepositoryWriter, mockNotifier, mockRepositoryTable)

	// Тестируем
	err := mockService.CreateTable()
	if err != nil {
		t.Fatalf("не удалось создать таблицу: %v", err)
	}

	if !mockRepositoryTable.Called {
		t.Fatalf("метод CreateTable не был вызван")
	}

	err = mockService.ProcessAndNotify("Иван", []string{"apple", "banana"}, 10.5)
	if err != nil {
		t.Fatalf("не удалось создать заказ: %v", err)
	}

	if !mockRepositoryWriter.Called {
		t.Fatalf("метод CreateOrder не был вызван")
	}

	expectedOrder := Order{
		Customer: "Иван",
		Products: []string{"apple", "banana"},
		Total: 10.5,
	}
	if !reflect.DeepEqual(mockRepositoryWriter.Order, expectedOrder) {
		t.Errorf("ожидали заказ: %+v, получили: %+v", expectedOrder, mockRepositoryWriter.Order)
	}

	if !mockNotifier.Called {
		t.Fatalf("метод Send не был вызван")
	}

	if mockNotifier.Customer != "Иван" {
		t.Errorf("ожидали клиента Иван, получили: %s", mockNotifier.Customer)
	}
}
