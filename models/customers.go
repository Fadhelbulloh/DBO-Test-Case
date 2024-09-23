package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/Fadhelbulloh/DBO-Test-Case/models/db"
	"github.com/Fadhelbulloh/DBO-Test-Case/params"
	"gorm.io/gorm"
)

type Customer struct {
	ID         uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT8;"`
	CustomerID string `json:"customer_id" gorm:"column:customer_id;type:VARCHAR(255);"`
	Name       string `json:"name" gorm:"column:name;type:VARCHAR(255);"`
	Email      string `json:"email" gorm:"column:email;type:VARCHAR(255);"`
	Phone      string `json:"phone" gorm:"column:phone;type:VARCHAR(255);"`
	CreatedAt  string `json:"created_at" gorm:"column:created_at;type:TIMESTAMPTZ;default:now();"`
	UpdatedAt  string `json:"updated_at" gorm:"column:updated_at;type:TIMESTAMPTZ;default:now();"`
}

// TableName overrides the default table name used by GORM
func (Customer) TableName() string {
	return "customers"
}

func GetAllCustomers(pagination params.Pagination, status, name, phone string) ([]Customer, error) {
	var customers []Customer
	dbConn := db.DB

	query := dbConn.Model(&Customer{})
	// Apply filters if provided
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if name != "" {
		query = query.Where("name = ?", name)
	}
	if phone != "" {
		query = query.Where("phone = ?", phone)
	}

	if pagination.Limit > 0 {
		query = query.Limit(pagination.Limit)
	}
	if pagination.Offset > 0 {
		query = query.Offset(pagination.Offset)
	}

	// Fetch data from DB
	if err := query.Find(&customers).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Printf("Error fetching customers: %v", err)
		return nil, err
	}

	return customers, nil
}

// GetCustomerByID retrieves a customer by ID
func GetCustomerByID(id string) (Customer, error) {
	var customer Customer

	// Retrieve customer by ID
	if err := db.DB.Where("customer_id = ? ", id).First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		}
		return customer, fmt.Errorf("error finding customer with ID %s: %v", id, err)
	}

	return customer, nil
}

// InsertCustomer inserts a new customer into the database
func InsertCustomer(customer Customer) error {
	if err := db.DB.Create(customer).Error; err != nil {
		return fmt.Errorf("error inserting customer: %v", err)
	}
	return nil
}

// DeleteCustomer deletes a customer from the database
func DeleteCustomer(id string) error {
	// Delete the customer where ID match the condition
	if err := db.DB.Where("customer_id = ? ", id).Delete(&Customer{}).Error; err != nil {
		return fmt.Errorf("error deleting customer with ID %s: %v", id, err)
	}
	return nil
}

// UpdateCustomer updates an existing customer in the database
func UpdateCustomer(id string, updatedData Customer) error {
	if err := db.DB.Model(&Customer{}).Where("customer_id = ?", id).Updates(updatedData).Error; err != nil {
		return fmt.Errorf("error updating customer with ID %s: %v", id, err)
	}
	return nil
}
