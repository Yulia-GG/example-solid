package sqlite3

import (
    //"example-solid/internal/repository/model"
    "fmt"
    "database/sql"
)

type SQLiteRepo struct {
db *sql.DB
}

// создадим конструктор
func NewSQLiteRepo(db *sql.DB) *SQLiteRepo {
   return &SQLiteRepo{
	db: db,
   }
}

func (s *SQLiteRepo) CreateOrder(customer string, products []string, total float64) error {
    _, err := s.db.Exec(
        "INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
        customer, fmt.Sprintf("%v", products), total, "pending",
    )
    if err != nil {
        return err
    }
      return nil
}

func (s *SQLiteRepo) CreateTable() error {
        _, err := s.db.Exec(`
    CREATE TABLE IF NOT EXISTS orders (
       id INTEGER PRIMARY KEY AUTOINCREMENT,
        customer TEXT NOT NULL,
        products TEXT NOT NULL,
        total REAL NOT NULL,
        status TEXT NOT NULL
    )`)
     if err != nil {
        return err
   }
      return nil
}
