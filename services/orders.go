package services

import (
	"github.com/Fadhelbulloh/DBO-Test-Case/models"
	"github.com/Fadhelbulloh/DBO-Test-Case/params"
	"github.com/google/uuid"
)

// GetOrders handles logic of paginated list of orders
func GetOrders(pagination params.Pagination, customerId, status string) ([]models.Order, error) {
	orders, err := models.GetOrders(pagination, customerId, status)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// GetOrderDetail handles logic of getting a single order
func GetOrderDetail(id string) (models.Order, error) {
	order, err := models.GetOrderDetail(id)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

// AddOrder handles adding a new order
func AddOrder(orderData models.Order) error {
	orderData.OrderNumber = uuid.New().String()
	if err := models.AddOrder(orderData); err != nil {
		return err
	}
	return nil
}

// UpdateOrder handles updating an existing order
func UpdateOrder(orderData models.Order, id string) error {
	if err := models.UpdateOrder(orderData, id); err != nil {
		return err
	}
	return nil
}

// DeleteOrder handles deleting an existing order
func DeleteOrder(id string) error {
	if err := models.DeleteOrder(id); err != nil {
		return err
	}
	return nil
}
