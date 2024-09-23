package services

import (
	"fmt"

	"github.com/Fadhelbulloh/DBO-Test-Case/models"
	"github.com/Fadhelbulloh/DBO-Test-Case/params"
	"github.com/google/uuid"
)

// AddCustomer handles logic of adding a new customer
func AddCustomer(custData models.Customer) error {
	custData.CustomerID = uuid.New().String()
	// Insert the customer
	if err := models.InsertCustomer(custData); err != nil {
		return err
	}
	return nil
}

// UpdateCustomer handles logic of updating an existing customer
func UpdateCustomer(custData models.Customer, id string) error {
	data, err := models.GetCustomerByID(id)
	if err != nil {
		return err
	}

	if data.ID == 0 {
		return fmt.Errorf("customer not found")
	}

	// Update the customer
	if err := models.UpdateCustomer(id, custData); err != nil {
		return err
	}
	return nil
}

// DeleteCustomer handles logic of deleting an existing customer
func DeleteCustomer(id string) error {
	data, err := models.GetCustomerByID(id)
	if err != nil {
		return err
	}

	if data.ID == 0 {
		return fmt.Errorf("customer not found")
	}

	// Delete the customer
	if err := models.DeleteCustomer(id); err != nil {
		return err
	}
	return nil
}

// GetCustomers handles logic of paginated list of customers
func GetCustomers(pagination params.Pagination, email, name, phone string) ([]models.Customer, error) {
	// Get customers
	customers, err := models.GetAllCustomers(pagination, email, name, phone)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// GetCustomerDetail handles logic of getting a single customer
func GetCustomerDetail(id string) (models.Customer, error) {
	// Get customer by ID
	customer, err := models.GetCustomerByID(id)
	if err != nil {
		return models.Customer{}, err
	}
	return customer, nil
}
