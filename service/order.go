package service

import (
	"time"
	"usecase_test/model"
	"usecase_test/repository"
)

type OrderService interface {
	FindAll() ([]model.OrderAllData, error)
	FindByID(id int) (model.OrderAllData, error)
	Create(order model.CreateOrderRequest) error
	// Update(order model.Order) (model.Order, error)
	// Delete(id int) (model.Order, error)
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{orderRepository}
}

func (s *orderService) FindAll() ([]model.OrderAllData, error) {
	orders, err := s.orderRepository.FindAll()
	if err != nil {
		return orders, err
	}

	if len(orders) > 0 {
		for i, order := range orders {
			orderDetails, err := s.orderRepository.FindOrderDetailByOrderID(int(order.ID))
			if err != nil {
				return orders, err
			}
			order.OrderDetails = orderDetails

			orderPayments, err := s.orderRepository.FindOrderPaymentByOrderID(int(order.ID))
			if err != nil {
				return orders, err
			}
			order.OrderPayments = orderPayments

			orders[i] = order
		}
	}

	return orders, nil
}

func (s *orderService) FindByID(id int) (model.OrderAllData, error) {
	order, err := s.orderRepository.FindByID(id)
	if err != nil {
		return order, err
	}

	orderDetails, err := s.orderRepository.FindOrderDetailByOrderID(id)
	if err != nil {
		return order, err
	}
	order.OrderDetails = orderDetails

	orderPayments, err := s.orderRepository.FindOrderPaymentByOrderID(id)
	if err != nil {
		return order, err
	}
	order.OrderPayments = orderPayments

	return order, nil
}

func (s *orderService) Create(order model.CreateOrderRequest) error {
	newOrder := model.Order{
		Date:              time.Now().Format("2006-01-02"),
		CustomerID:        order.CustomerID,
		CustomerAddressID: order.CustomerAddressID,
		Total:             order.TotalPrice,
	}

	createdOrder, err := s.orderRepository.Create(newOrder)
	if err != nil {
		return err
	}

	for _, orderDetail := range order.OrderDetails {
		newOrderDetail := model.OrderDetails{
			OrderID:   createdOrder.ID,
			ProductID: orderDetail.ProductID,
			Quantity:  orderDetail.Quantity,
			SubTotal:  orderDetail.SubTotal,
		}

		err = s.orderRepository.CreateOrderDetail(newOrderDetail)
		if err != nil {
			return err
		}
	}

	for _, orderPayment := range order.OrderPayments {
		newOrderPayment := model.OrderPayments{
			OrderID:         createdOrder.ID,
			PaymentMethodID: orderPayment.PaymentMethodID,
			Info:            orderPayment.Info,
		}

		err = s.orderRepository.CreateOrderPayment(newOrderPayment)
		if err != nil {
			return err
		}
	}

	return nil
}
