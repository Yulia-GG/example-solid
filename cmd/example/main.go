package main

import (
	"database/sql"
	"example-solid/internal/notifier"
	"example-solid/internal/repository/sqlite3"
	"example-solid/internal/service"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// открываем базу  данных
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // закрываем б.д. после открытия

	// создаем репозиторий (первый вариант)
	// repo := &sqlite3.SQLiteRepo{
	//	db: db,
	//}

	// (второй вариант - через конструктор)
	repo := sqlite3.NewSQLiteRepo(db)

	// создание SMS системы
	SMSService := service.NewOrderService(
		repo,
		&notifier.SMSSender{},
		repo,
	)

	// создание таблицы
	err = SMSService.CreateTable()
	if err != nil {
		log.Fatal(err)
	}

	// Пример использования
	err = SMSService.ProcessAndNotify("Иван", []string{"apple", "banana"}, 10.5)
	if err != nil {
		log.Fatal(err)
	}

	// создание Email системы
EmailService := service.NewOrderService(
		repo,
		&notifier.EmailSender{},
		repo,
	)

// Пример использования
	err = EmailService.ProcessAndNotify("Николай", []string{"apple", "mango"}, 15.5)
	if err != nil {
		log.Fatal(err)
	}

}
