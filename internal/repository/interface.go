package repository

// import "example-solid/internal/repository/model"

// Интерфейс создания заказа для всех типов баз данных
type RepositoryWriter interface {
CreateOrder(customer string, products []string, total float64) error
}

// Интерфейс для создания таблицы
type RepositoryTable interface {
CreateTable() error
}

// Интерфейс для отправки уведомлений
type Notifier interface {
Send(customer string)
}