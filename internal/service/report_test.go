package service

import(
"testing"
)

type MockRepositoryWriter struct {}

func (m *MockRepositoryWriter) CreateOrder(customer string, products []string, total float64) error {
return nil
}

type MockNotifier struct {}

func (n *MockNotifier)  Send(customer string) {}

type MockRepositoryTable struct {}

func (t *MockRepositoryTable) CreateTable() error {
return nil
}

func TestOrderService(t *testing.T) {

// Передаем моки в сервис
service := NewOrderService(&MockRepositoryWriter{}, &MockNotifier{}, &MockRepositoryTable{})

// Тестируем
err := service.CreateTable()
if err != nil {
		t.Fatalf("не удалось создать таблицу: %v", err)
	}
err = service.ProcessAndNotify("Иван", []string{"apple", "banana"}, 10.5)
if err != nil {
		t.Fatalf("не удалось создать заказ: %v", err)
	}


}
