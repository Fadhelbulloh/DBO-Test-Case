package models

import (
	"errors"
	"log"
	"time"

	"github.com/Fadhelbulloh/DBO-Test-Case/models/db"
	"github.com/Fadhelbulloh/DBO-Test-Case/params"
	"gorm.io/gorm"
)

type Order struct {
	ID            uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT8;"`
	OrderNumber   string    `json:"order_number" gorm:"column:order_number;type:VARCHAR(255);"`
	CustomerID    uint      `json:"customer_id" gorm:"column:customer_id;type:INT8;"`
	TotalAmount   float64   `json:"total_amount" gorm:"column:total_amount;type:FLOAT8;"`
	Status        string    `json:"status" gorm:"column:status;type:VARCHAR(255);"`
	PaymentMethod string    `json:"payment_method" gorm:"column:payment_method;type:VARCHAR(255);"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at;type:TIMESTAMPTZ;default:now();"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at;type:TIMESTAMPTZ;default:now();"`
}

// TableName overrides the default table name used by GORM
func (Order) TableName() string {
	return "orders"
}

// GetOrders handles logic of paginated list of orders
func GetOrders(pagination params.Pagination, customerId, status string) ([]Order, error) {
	var orders []Order
	dbConn := db.DB

	query := dbConn.Model(&Order{})
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if customerId != "" {
		query = query.Where("customer_id = ?", customerId)
	}

	if pagination.Limit > 0 {
		query = query.Limit(pagination.Limit)
	}

	if pagination.Offset > 0 {
		query = query.Offset(pagination.Offset)
	}
	// Fetch data from DB
	if err := query.Find(&orders).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Printf("Error fetching customers: %v", err)
		return nil, err
	}

	return orders, nil
}

// GetOrderDetail handles logic of getting a single order
func GetOrderDetail(id string) (Order, error) {
	var order Order
	dbConn := db.DB
	if err := dbConn.Where("order_number = ?", id).First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, nil
		}
		log.Printf("Error fetching order: %v", err)
		return order, err
	}
	return order, nil
}

// AddOrder handles adding a new order
func AddOrder(order Order) error {
	dbConn := db.DB
	if err := dbConn.Create(&order).Error; err != nil {
		log.Printf("Error creating order: %v", err)
		return err
	}
	return nil
}

// UpdateOrder handles updating an existing order
func UpdateOrder(order Order, id string) error {
	dbConn := db.DB
	if err := dbConn.Where("order_number = ?", id).Updates(&order).Error; err != nil {
		log.Printf("Error updating order: %v", err)
		return err
	}
	return nil
}

// DeleteOrder handles deleting an existing order
func DeleteOrder(id string) error {
	dbConn := db.DB
	if err := dbConn.Where("order_number = ?", id).Delete(&Order{}).Error; err != nil {
		log.Printf("Error deleting order: %v", err)
		return err
	}
	return nil
}
