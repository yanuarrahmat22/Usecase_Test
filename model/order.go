package model

type Order struct {
	ID                uint64 `gorm:"id" `
	Date              string `gorm:"date" `
	Total             uint64 `gorm:"total" `
	CustomerID        uint64 `gorm:"customer_id" `
	CustomerAddressID uint64 `gorm:"customer_address_id" `
}

type OrderDetail struct {
	OrderID     uint64 `gorm:"order_id" json:"order_id"`
	ProductID   uint64 `gorm:"product_id" json:"product_id"`
	ProductName string `gorm:"product_name" json:"product_name"`
	Price       uint64 `gorm:"price" json:"price"`
	Quantity    uint64 `gorm:"quantity" json:"quantity"`
	SubTotal    uint64 `gorm:"sub_total" json:"sub_total"`
}

type OrderPayment struct {
	OrderID           uint64 `gorm:"order_id" json:"order_id"`
	PaymentMethodID   uint64 `gorm:"payment_method_id" json:"payment_method_id"`
	PaymentMethodName string `gorm:"payment_method_name" json:"payment_method_name"`
	IsActive          int    `gorm:"is_active" json:"is_active"`
	Info              string `gorm:"info" json:"info"`
}

type OrderAllData struct {
	ID                uint64         `gorm:"id" json:"id"`
	Date              string         `gorm:"date" json:"date"`
	TotalPrice        uint64         `gorm:"total" json:"total_price"`
	CustomerID        uint64         `gorm:"customer_id" json:"customer_id"`
	CustomerName      string         `gorm:"customer_name" json:"customer_name"`
	CustomerAddressID uint64         `gorm:"customer_address_id" json:"customer_address_id"`
	CustomerAddress   string         `gorm:"customer_address" json:"customer_address"`
	OrderDetails      []OrderDetail  `gorm:"-" json:"order_details"`  // Exclude this field from raw query
	OrderPayments     []OrderPayment `gorm:"-" json:"order_payments"` // Exclude this field from raw query
}

type CreateOrderRequest struct {
	CustomerID        uint64 `json:"customer_id" validate:"required"`
	CustomerAddressID uint64 `json:"customer_address_id" validate:"required"`
	TotalPrice        uint64 `json:"total_price" validate:"required"`
	OrderDetails      []struct {
		ProductID uint64 `json:"product_id" validate:"required"`
		Quantity  uint64 `json:"quantity" validate:"required"`
		SubTotal  uint64 `json:"sub_total" validate:"required"`
	} `json:"order_details" validate:"required"`
	OrderPayments []struct {
		PaymentMethodID uint64 `json:"payment_method_id" validate:"required"`
		Info            string `json:"info" validate:"required"`
	} `json:"order_payments" validate:"required"`
}

type OrderDetails struct {
	OrderID   uint64 
	ProductID uint64
	Quantity  uint64
	SubTotal  uint64
}

type OrderPayments struct {
	OrderID         uint64
	PaymentMethodID uint64
	Info            string
}