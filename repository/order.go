package repository

import (
	"log"
	"usecase_test/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAll() ([]model.OrderAllData, error)
	FindByID(id int) (model.OrderAllData, error)
	FindOrderDetailByOrderID(id int) ([]model.OrderDetail, error)
	FindOrderPaymentByOrderID(id int) ([]model.OrderPayment, error)
	Create(order model.Order) (model.Order, error)
	CreateOrderDetail(orderDetail model.OrderDetails) error
	CreateOrderPayment(orderPayment model.OrderPayments) error
	// Update(order model.Order) (model.Order, error)
	// Delete(id int) (model.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) FindAll() ([]model.OrderAllData, error) {
	var orders []model.OrderAllData
	err := r.db.Raw(`
		SELECT  o.id, o.date, o.total, o.customer_id, c.customer_name, o.customer_address_id, ca.address_name as customer_address
		FROM orders o 
		JOIN customer c ON c.id = o.customer_id
		JOIN customer_address ca ON ca.id = o.customer_address_id;
	`).Scan(&orders).Error
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (r *orderRepository) FindByID(id int) (model.OrderAllData, error) {
	var order model.OrderAllData
	err := r.db.Raw(`
		SELECT  o.id, o.date, o.total, o.customer_id, c.customer_name, o.customer_address_id, ca.address_name as customer_address
		FROM orders o 
		JOIN customer c ON c.id = o.customer_id
		JOIN customer_address ca ON ca.id = o.customer_address_id
		WHERE o.id = ?;
	`, id).Scan(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *orderRepository) FindOrderDetailByOrderID(id int) ([]model.OrderDetail, error) {
	var orderDetails []model.OrderDetail
	err := r.db.Raw(`
		SELECT od.order_id, od.product_id, p.name as product_name, p.price, od.quantity, od.sub_total
		FROM order_details od
		JOIN product p ON p.id = od.product_id
		WHERE od.order_id = ?;
	`, id).Scan(&orderDetails).Error
	if err != nil {
		return orderDetails, err
	}

	return orderDetails, nil
}

func (r *orderRepository) FindOrderPaymentByOrderID(id int) ([]model.OrderPayment, error) {
	var orderPayments []model.OrderPayment
	err := r.db.Raw(`
		SELECT op.order_id, op.payment_method_id, pm.name as payment_method_name, op.info, pm.is_active
		FROM order_payments op 
		JOIN payment_method pm ON pm.id = op.payment_method_id
		WHERE op.order_id = ?;
	`, id).Scan(&orderPayments).Error
	if err != nil {
		return orderPayments, err
	}

	return orderPayments, nil
}

func (r *orderRepository) CreateOrderDetail(orderDetail model.OrderDetails) error {
	err := r.db.Create(&orderDetail).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) CreateOrderPayment(orderPayment model.OrderPayments) error {
	err := r.db.Create(&orderPayment).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) Create(order model.Order) (model.Order, error) {
	var lastOrderID uint64
	if err := r.db.Model(&order).Select("id").Order("id desc").Limit(1).Scan(&lastOrderID).Error; err != nil {
		log.Fatal(err)
	}

	order.ID = uint64(lastOrderID + 1)

	result := r.db.Create(&order)
	if result.Error != nil {
		return order, result.Error
	}

	return order, nil
}
