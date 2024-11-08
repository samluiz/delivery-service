package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/samluiz/delivery-service/internal/delivery"
	"github.com/stretchr/testify/assert"
)

type MockDeliveryService struct {
	CreateDeliveryFn      func(req *delivery.CreateDeliveryRequest) (*delivery.DeliveryResponse, error)
	GetDeliveryFn         func(id int) (*delivery.DeliveryResponse, error)
	GetDeliveriesFn       func(city string) ([]*delivery.DeliveryResponse, error)
	UpdateDeliveryFn      func(req *delivery.UpdateDeliveryRequest, id int) (*delivery.DeliveryResponse, error)
	DeleteDeliveryFn      func(id int) error
	DeleteAllDeliveriesFn func() error
}

func (m MockDeliveryService) CreateDelivery(req *delivery.CreateDeliveryRequest) (*delivery.DeliveryResponse, error) {
	return m.CreateDeliveryFn(req)
}

func (m MockDeliveryService) GetDelivery(id int) (*delivery.DeliveryResponse, error) {
	return m.GetDeliveryFn(id)
}

func (m MockDeliveryService) GetDeliveries(city string) ([]*delivery.DeliveryResponse, error) {
	return m.GetDeliveriesFn(city)
}

func (m MockDeliveryService) UpdateDelivery(req *delivery.UpdateDeliveryRequest, id int) (*delivery.DeliveryResponse, error) {
	return m.UpdateDeliveryFn(req, id)
}

func (m MockDeliveryService) DeleteDelivery(id int) error {
	return m.DeleteDeliveryFn(id)
}

func (m MockDeliveryService) DeleteAllDeliveries() error {
	return m.DeleteAllDeliveriesFn()
}

func TestHandleCreateDelivery(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int
		expectedError  error
	}{
		{
			name: "valid request",
			requestBody: &delivery.CreateDeliveryRequest{
				Cliente:     "Cliente A",
				Peso:        10.5,
				Endereco:    "Rua A",
				Logradouro:  "Logradouro A",
				Complemento: "Complemento A",
				Numero:      "123",
				Bairro:      "Bairro A",
				Cidade:      "Cidade A",
				Estado:      "Estado A",
				Pais:        "Brasil",
				Latitude:    45.5,
				Longitude:   12.5,
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "invalid request body",
			requestBody: struct {
				Cliente string `json:"cliente"`
			}{},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deliveryServiceMock := MockDeliveryService{
				CreateDeliveryFn: func(req *delivery.CreateDeliveryRequest) (*delivery.DeliveryResponse, error) {
					if tt.expectedError != nil {
						return nil, tt.expectedError
					}
					return &delivery.DeliveryResponse{ID: 1, Cliente: req.Cliente}, nil
				},
			}
			handler := NewDeliveryHandler(deliveryServiceMock)

			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("POST", "/deliveries", bytes.NewReader(body))
			w := httptest.NewRecorder()

			handler.HandleCreateDelivery(w, req)

			res := w.Result()
			assert.Equal(t, tt.expectedStatus, res.StatusCode)
		})
	}
}

func TestHandleGetDelivery(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		expectedStatus int
		expectedError  error
	}{
		{
			name:           "valid request",
			id:             "1",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid id",
			id:             "invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "not found",
			id:             "2",
			expectedStatus: http.StatusNotFound,
			expectedError:  delivery.ErrDeliveryNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deliveryServiceMock := MockDeliveryService{
				GetDeliveryFn: func(id int) (*delivery.DeliveryResponse, error) {
					if tt.expectedError != nil {
						return nil, tt.expectedError
					}
					return &delivery.DeliveryResponse{ID: id, Cliente: "Client A"}, nil
				},
			}

			handler := NewDeliveryHandler(deliveryServiceMock)

			mux := http.NewServeMux()
			mux.HandleFunc("/deliveries/{id}", handler.HandleGetDelivery)

			req := httptest.NewRequest("GET", fmt.Sprintf("/deliveries/%s", tt.id), nil)
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, req)

			res := w.Result()
			assert.Equal(t, tt.expectedStatus, res.StatusCode)
		})
	}
}

func TestHandleGetDeliveries(t *testing.T) {
	tests := []struct {
		name           string
		city           string
		expectedStatus int
		expectedError  error
	}{
		{
			name:           "valid city",
			city:           "Cidade A",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "no city",
			city:           "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deliveryServiceMock := MockDeliveryService{
				GetDeliveriesFn: func(city string) ([]*delivery.DeliveryResponse, error) {
					if tt.expectedError != nil {
						return nil, tt.expectedError
					}
					return []*delivery.DeliveryResponse{
						{ID: 1, Cliente: "Client A", Cidade: city},
					}, nil
				},
			}
			handler := NewDeliveryHandler(deliveryServiceMock)

			cityEncoded := url.QueryEscape(tt.city)
			var req *http.Request

			if tt.city == "" {
				req = httptest.NewRequest("GET", "/deliveries", nil)
			} else {
				req = httptest.NewRequest("GET", fmt.Sprintf("/deliveries?city=%s", cityEncoded), nil)
			}

			w := httptest.NewRecorder()

			handler.HandleGetDeliveries(w, req)

			res := w.Result()
			assert.Equal(t, tt.expectedStatus, res.StatusCode)
		})
	}
}

func TestHandleUpdateDelivery(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		requestBody    interface{}
		expectedStatus int
		expectedError  error
	}{
		{
			name: "valid request",
			id:   "1",
			requestBody: &delivery.UpdateDeliveryRequest{
				Peso:        15.5,
				Endereco:    "Rua A",
				Logradouro:  "Logradouro A",
				Complemento: "Complemento A",
				Numero:      "123",
				Bairro:      "Bairro A",
				Cidade:      "Cidade A",
				Estado:      "Estado A",
				Pais:        "Brasil",
				Latitude:    45.5,
				Longitude:   12.5,
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid id",
			id:             "invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "delivery not found",
			id:             "2",
			requestBody:    &delivery.UpdateDeliveryRequest{
				Peso:        15.5,
				Endereco:    "Rua B",
				Logradouro:  "Logradouro A",
				Complemento: "Complemento A",
				Numero:      "123",
				Bairro:      "Bairro A",
				Cidade:      "Cidade A",
				Estado:      "Estado A",
				Pais:        "Brasil",
				Latitude:    45.5,
				Longitude:   12.5,
			},
			expectedStatus: http.StatusNotFound,
			expectedError:  delivery.ErrDeliveryNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deliveryServiceMock := MockDeliveryService{
				UpdateDeliveryFn: func(req *delivery.UpdateDeliveryRequest, id int) (*delivery.DeliveryResponse, error) {
					if tt.expectedError != nil {
						return nil, tt.expectedError
					}
					return &delivery.DeliveryResponse{ID: id, Cliente: "Client A"}, nil
				},
			}
			handler := NewDeliveryHandler(deliveryServiceMock)

			mux := http.NewServeMux()
			mux.HandleFunc("/deliveries/{id}", handler.HandleUpdateDelivery)

			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("PUT", fmt.Sprintf("/deliveries/%s", tt.id), bytes.NewReader(body))

			w := httptest.NewRecorder()

			mux.ServeHTTP(w, req)

			res := w.Result()
			assert.Equal(t, tt.expectedStatus, res.StatusCode)
		})
	}
}

func TestHandleDeleteDelivery(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		expectedStatus int
		expectedError  error
	}{
		{
			name:           "valid request",
			id:             "1",
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "invalid id",
			id:             "invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "delivery not found",
			id:             "2",
			expectedStatus: http.StatusNotFound,
			expectedError:  delivery.ErrDeliveryNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deliveryServiceMock := MockDeliveryService{
				DeleteDeliveryFn: func(id int) error {
					if tt.expectedError != nil {
						return tt.expectedError
					}
					return nil
				},
			}
			handler := NewDeliveryHandler(deliveryServiceMock)

			mux := http.NewServeMux()
			mux.HandleFunc("/deliveries/{id}", handler.HandleDeleteDelivery)

			req := httptest.NewRequest("DELETE", fmt.Sprintf("/deliveries/%s", tt.id), nil)

			w := httptest.NewRecorder()

			mux.ServeHTTP(w, req)

			res := w.Result()
			assert.Equal(t, tt.expectedStatus, res.StatusCode)
		})
	}
}


func TestHandleDeleteAllDeliveries(t *testing.T) {
	deliveryServiceMock := MockDeliveryService{
		DeleteAllDeliveriesFn: func() error {
			return nil
		},
	}
	handler := NewDeliveryHandler(deliveryServiceMock)

	req := httptest.NewRequest("DELETE", "/deliveries", nil)
	w := httptest.NewRecorder()

	handler.HandleDeleteAllDeliveries(w, req)

	res := w.Result()
	assert.Equal(t, http.StatusNoContent, res.StatusCode)
}
