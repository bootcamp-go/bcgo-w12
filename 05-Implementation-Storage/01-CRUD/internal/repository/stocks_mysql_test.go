package repository_test

import (
	"database/sql"
	"implementacion-storage/crud/internal/repository"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func init() {
	// set-up database
	cfg := mysql.Config{
		User:	"stocks_test_user",
		Passwd: "StocksTestUser123!",
		Net:	"tcp",
		Addr:	"localhost:3306",
		DBName:	"stocks_test_db",
		ParseTime: true,
	}
	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}

func TestStockMySQL_Find(t *testing.T) {
	t.Run("success - find some data", func(t *testing.T) {
		// arrange
		// - database
		db, err := sql.Open("txdb", "success_find_some_data")
		require.NoError(t, err)
		// - database: clean-up / rollback
		defer func(db *sql.DB) {
			// manual rollback since txdb doesn't support rollback for DDL
			// delete all stocks
			db.Exec("DELETE FROM stocks")
			// reset auto-increment
			db.Exec("ALTER TABLE stocks AUTO_INCREMENT = 1")

			// by using DDL, we can't rollback the transaction (forced manually)
			db.Close()
		}(db)
		// - database: set-up
		err = func(db *sql.DB) (e error) {
			// insert 2 stocks
			stmt, e := db.Prepare("INSERT INTO stocks (name, symbol, price) VALUES (?, ?, ?)")
			if e != nil {
				return
			}
			defer stmt.Close()

			_, e = stmt.Exec("Apple Inc.", "AAPL", 123.45)
			if e != nil {
				return
			}
			_, e = stmt.Exec("Microsoft Corporation", "MSFT", 67.89)
			if e != nil {
				return
			}
			return
		}(db)
		require.NoError(t, err)
		// - repository
		rp := repository.NewStockMySQL(db)

		// act
		stocks, err := rp.Find()

		// assert
		require.NoError(t, err)
		require.Len(t, stocks, 2)
		require.Equal(t, 1, stocks[0].ID)
		require.Equal(t, "Apple Inc.", stocks[0].Name)
		require.Equal(t, "AAPL", stocks[0].Symbol)
		require.Equal(t, 123.45, stocks[0].Price)
		require.Equal(t, 2, stocks[1].ID)
		require.Equal(t, "Microsoft Corporation", stocks[1].Name)
		require.Equal(t, "MSFT", stocks[1].Symbol)
		require.Equal(t, 67.89, stocks[1].Price)
	})

	t.Run("success - find no data", func(t *testing.T) {
		// arrange
		// - database
		db, err := sql.Open("txdb", "success_find_no_data")
		require.NoError(t, err)
		// - database: clean-up / rollback
		defer db.Close()
		// - repository
		rp := repository.NewStockMySQL(db)

		// act
		stocks, err := rp.Find()

		// assert
		require.NoError(t, err)
		require.Len(t, stocks, 0)
		require.Nil(t, stocks)
		// require.Equal(t, []internal.Stock(nil), stocks)
	})
}