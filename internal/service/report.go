package service

import (
    "example-solid/internal/repository"
    "fmt"
)

// Интерфейс для отправки уведомлений
type Notifier interface {
Send(customer string) error
}

// главный сервис, объединяющий компоненты
type OrderService struct {
repositoryWriter repository.RepositoryWriter // Встраивание интерфейса
notifier Notifier // Встраивание интерфейса
repositoryTable repository.RepositoryTable  // Встраивание интерфейса
}

// создаем таблицу
func (s *OrderService) CreateTable() error {
    err := s.repositoryTable.CreateTable()
    if err != nil {
    return fmt.Errorf("ошибка создания таблицы: %w", err)
	}
    return nil
}

func (s *OrderService) ProcessAndNotify(customer string, products []string, total float64) error  {
    // Создаем заказ
    err := s.repositoryWriter.CreateOrder(customer, products, total)
    if err != nil {
    return fmt.Errorf("ошибка создания заказа: %w", err)
	}

    // Отправляем уведомление
    err = s.notifier.Send(customer)
    if err != nil {
    return fmt.Errorf("ошибка отправления уведомления: %w", err)
	}

	return nil
}

//func (s *OrderService) ProcessAndNotify(customer string, products []string, total float64) error  {
    // Создаем заказ
    //err := s.repositoryWriter.CreateOrder(customer, products, total)
    //if err != nil {
    //return fmt.Errorf("ошибка создания заказа: %w", err)
    //}
    // Отправляем уведомление
    //s.notifier.Send(customer)
    //return nil
//}

// создадим конструктор
func NewOrderService(repositoryWriter repository.RepositoryWriter, notifier Notifier, repositoryTable repository.RepositoryTable ) *OrderService {
return &OrderService{
repositoryWriter: repositoryWriter,
notifier: notifier,
repositoryTable: repositoryTable,
}
}
