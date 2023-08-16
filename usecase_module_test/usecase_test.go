package usecasemodule_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"usecase_test/handler"
	"usecase_test/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockOrderService struct {
	mock.Mock
}

func (m *mockOrderService) FindAll() ([]model.OrderAllData, error) {
	args := m.Called()
	return args.Get(0).([]model.OrderAllData), args.Error(1)
}

func (m *mockOrderService) FindByID(id int) (model.OrderAllData, error) {
	args := m.Called(id)
	return args.Get(0).(model.OrderAllData), args.Error(1)
}

func (m *mockOrderService) Create(order model.CreateOrderRequest) error {
	args := m.Called(order)
	return args.Error(0)
}

func TestFindAll(t *testing.T) {
	mockService := new(mockOrderService)
	mockService.On("FindAll").Return([]model.OrderAllData{}, nil)

	handler := handler.NewOrderHandler(mockService)
	router := gin.Default()
	router.GET("/orders", handler.FindAll)

	req, _ := http.NewRequest("GET", "/orders", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

func TestFindByID(t *testing.T) {
	mockService := new(mockOrderService)
	expectedOrder := model.OrderAllData{ID: 1, CustomerName: "John Doe", TotalPrice: 100000}
	mockService.On("FindByID", 1).Return(expectedOrder, nil)

	handler := handler.NewOrderHandler(mockService)
	router := gin.Default()
	router.GET("/orders/:id", handler.FindByID)

	req, _ := http.NewRequest("GET", "/orders/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	mockService := new(mockOrderService)
	mockService.On("Create", mock.Anything).Return(nil)

	handler := handler.NewOrderHandler(mockService)
	router := gin.Default()
	router.POST("/orders", handler.Create)

	// Simulate a valid JSON request
	orderJSON := `
		{
			"customer_id" : 1,
			"customer_address_id" : 1,
			"total_price" : 25000,
			"order_details": [
				{
					"product_id": 1,
					"quantity": 1,
					"sub_total": 25000
				}
			],
			"order_payments": [
				{
					"payment_method_id": 1,
					"info": "gini aja"
				}
			]
		}
	`
	req, _ := http.NewRequest("POST", "/orders", strings.NewReader(orderJSON))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	mockService.AssertExpectations(t)
}

func TestCreateInvalidData(t *testing.T) {
	mockService := new(mockOrderService)

	handler := handler.NewOrderHandler(mockService)
	router := gin.Default()
	router.POST("/orders", handler.Create)

	// Sending invalid JSON data
	req, _ := http.NewRequest("POST", "/orders", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	mockService.AssertNotCalled(t, "Create")
}

func TestFindAllError(t *testing.T) {
	mockService := new(mockOrderService)
	mockService.On("FindAll").Return(nil, errors.New("error"))

	handler := handler.NewOrderHandler(mockService)
	router := gin.Default()
	router.GET("/orders", handler.FindAll)

	req, _ := http.NewRequest("GET", "/orders", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	mockService.AssertExpectations(t)
}
