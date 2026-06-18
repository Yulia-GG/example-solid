package main

import (
	"example-solid/internal/repository/sqlite3"
    "example-solid/internal/notifier"
    "example-solid/internal/service"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// открываем базу  данных
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	// создаем репозиторий (первый вариант)
	// repo := &sqlite3.SQLiteRepo{
	//	db: db,
	//}

	// (второй вариант - через конструктор)
	repo := sqlite3.NewSQLiteRepo(db)

	// создание системы
	service := service.NewOrderService(
		repo,
		&notifier.EmailSender{},
		repo,
	)

	// создание таблицы
	err = service.CreateTable()
	if err != nil {
		log.Fatal(err)
	}

	// Пример использования
	err = service.ProcessAndNotify("Иван", []string{"apple", "banana"}, 10.5)
	if err != nil {
		log.Fatal(err)
	}
}
