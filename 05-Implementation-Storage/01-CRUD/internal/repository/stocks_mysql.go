package repository

import (
	"database/sql"
	"errors"
	"implementacion-storage/crud/internal"

	"github.com/go-sql-driver/mysql"
)

var (
	// ErrStockRowsAffected is an error that is returned when the number of rows affected is not 1
	ErrStockRowsAffected = errors.New("number of rows affected is not the expected")
)

func NewStockMySQL(db *sql.DB) *StockMySQL {
	return &StockMySQL{db: db}
}

// type StockScheme struct {
// 	ID        sql.NullInt64
// 	Name      sql.NullString
// 	Symbol    sql.NullString
// 	Price     sql.NullFloat64
// 	CreatedAt sql.NullTime
// 	UpdatedAt sql.NullTime
// }

// StockMySQL is a struct that represents a MySQL implementation of the StockRepository
type StockMySQL struct {
	// db is a pool of database connections
	db *sql.DB
}

// Find returns stocks that match the given criteria
func (s *StockMySQL) Find() (stocks []internal.Stock, err error) {
	// execute the query
	query := "SELECT id, name, symbol, price, created_at, updated_at FROM stocks"
	rows, err := s.db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var stock internal.Stock
		
		// scan the result
		err = rows.Scan(&stock.ID, &stock.Name, &stock.Symbol, &stock.Price, &stock.CreatedAt, &stock.UpdatedAt)
		if err != nil {
			return
		}

		// append the stock to the slice
		stocks = append(stocks, stock)
	}

	// check for errors during the iteration
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// FindByID returns a stock that matches the given ID
func (s *StockMySQL) FindByID(id int) (stock internal.Stock, err error) {
	// execute the query
	query := "SELECT id, name, symbol, price, created_at, updated_at FROM stocks WHERE id = ?"
	row := s.db.QueryRow(query, id)
	if row.Err() != nil {
		err = row.Err()
		return
	}

	// scan the result
	err = row.Scan(&stock.ID, &stock.Name, &stock.Symbol, &stock.Price, &stock.CreatedAt, &stock.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrStockNotFound
			return
		}
		return
	}

	return
}

// Save saves the given stock
func (s *StockMySQL) Save(stock *internal.Stock) (err error) {
	// execute the query
	query := "INSERT INTO stocks (name, symbol, price) VALUES (?, ?, ?)"
	result, err := s.db.Exec(query, stock.Name, stock.Symbol, stock.Price)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrStockDuplicated
			default:
				// ...
			}
			return
		}

		return
	}

	// check result
	// - rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected != 1 {
		err = ErrStockRowsAffected
		return
	}
	// - last insert ID
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return
	}

	// update the ID
	stock.ID = int(lastInsertedID)

	return
}

// Update updates the given stock
func (s *StockMySQL) Update(stock *internal.Stock) (err error) {
	// execute the query
	query := "UPDATE stocks SET name = ?, symbol = ?, price = ? WHERE id = ?"
	result, err := s.db.Exec(query, stock.Name, stock.Symbol, stock.Price, stock.ID)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrStockDuplicated
			default:
				// ...
			}
			return
		}

		return
	}

	// check result
	// - rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected != 1 {
		err = ErrStockRowsAffected
		return
	}

	return
}

// Delete deletes the stock that matches the given ID
func (s *StockMySQL) Delete(id int) (err error) {
	// execute the query
	query := "DELETE FROM stocks WHERE id = ?"
	result, err := s.db.Exec(query, id)
	if err != nil {
		return
	}

	// check result
	// - rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected != 1 {
		switch rowsAffected {
		case 0:
			err = internal.ErrStockNotFound
		default:
			err = ErrStockRowsAffected
		}
		return
	}

	return
}
	
