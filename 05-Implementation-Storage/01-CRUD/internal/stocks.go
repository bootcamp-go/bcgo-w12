package internal

import (
	"errors"
	"time"
)

var (
	// ErrStockNotFound is an error that is returned when the stock is not found
	ErrStockNotFound = errors.New("stock not found")
	// ErrStockDuplicated is an error that is returned when the stock is duplicated
	ErrStockDuplicated = errors.New("stock duplicated")
)

// Stock is a struct that represents a stock
type Stock struct {
	// ID is the unique identifier of the stock
	ID        int
	// Name is the name of the stock
	Name      string
	// Symbol is the symbol of the stock
	// - unique
	Symbol    string
	// Price is the price of the stock
	Price     float64
	// CreatedAt is the timestamp when the stock was created
	CreatedAt time.Time
	// UpdatedAt is the timestamp when the stock was last updated
	UpdatedAt time.Time
}

// StockRepository is an interface that defines the methods to interact with the stocks
type StockRepository interface {
	// Read Operations
	// Find returns stocks that match the given criteria
	Find() ([]Stock, error)
	// FindByID returns a stock that matches the given ID
	FindByID(id int) (Stock, error)
	
	// Write Operations
	// Save saves the given stock
	Save(stock *Stock) error
	// Update updates the given stock
	Update(stock *Stock) error
	// Delete deletes the stock that matches the given ID
	Delete(id int) error
}

