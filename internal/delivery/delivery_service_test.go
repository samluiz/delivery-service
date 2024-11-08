package delivery

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDeliveryRepository struct {
	mock.Mock
}

func (m *MockDeliveryRepository) CreateDelivery(request *CreateDeliveryRequest) (*DeliveryResponse, error) {
	args := m.Called(request)
	return args.Get(0).(*DeliveryResponse), args.Error(1)
}

func (m *MockDeliveryRepository) GetDelivery(id int) (*DeliveryResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*DeliveryResponse), args.Error(1)
}

func (m *MockDeliveryRepository) GetDeliveries() ([]*DeliveryResponse, error) {
	args := m.Called()
	return args.Get(0).([]*DeliveryResponse), args.Error(1)
}

func (m *MockDeliveryRepository) GetDeliveriesByCity(city string) ([]*DeliveryResponse, error) {
	args := m.Called(city)
	return args.Get(0).([]*DeliveryResponse), args.Error(1)
}

func (m *MockDeliveryRepository) UpdateDelivery(request *UpdateDeliveryRequest, id int) (*DeliveryResponse, error) {
	args := m.Called(request, id)
	return args.Get(0).(*DeliveryResponse), args.Error(1)
}

func (m *MockDeliveryRepository) DeleteDelivery(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockDeliveryRepository) DeleteAllDeliveries() error {
	args := m.Called()
	return args.Error(0)
}

func TestCreateDelivery(t *testing.T) {
	mockRepo := new(MockDeliveryRepository)
	service := NewDeliveryService(mockRepo)

	request := &CreateDeliveryRequest{}
	expectedResponse := &DeliveryResponse{}

	mockRepo.On("CreateDelivery", request).Return(expectedResponse, nil)

	response, err := service.CreateDelivery(request)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
	mockRepo.AssertExpectations(t)
}

func TestGetDelivery(t *testing.T) {
	mockRepo := new(MockDeliveryRepository)
	service := NewDeliveryService(mockRepo)

	id := 1
	expectedResponse := &DeliveryResponse{}

	mockRepo.On("GetDelivery", id).Return(expectedResponse, nil)

	response, err := service.GetDelivery(id)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
	mockRepo.AssertExpectations(t)
}

func TestGetDeliveries(t *testing.T) {
	mockRepo := new(MockDeliveryRepository)
	service := NewDeliveryService(mockRepo)

	expectedResponse := []*DeliveryResponse{}
	mockRepo.On("GetDeliveries").Return(expectedResponse, nil)

	response, err := service.GetDeliveries("")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
	mockRepo.AssertExpectations(t)

	city := "City1"
	mockRepo.On("GetDeliveriesByCity", city).Return(expectedResponse, nil)

	response, err = service.GetDeliveries(city)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
	mockRepo.AssertExpectations(t)
}

func TestUpdateDelivery(t *testing.T) {
	mockRepo := new(MockDeliveryRepository)
	service := NewDeliveryService(mockRepo)

	request := &UpdateDeliveryRequest{}
	id := 1
	expectedResponse := &DeliveryResponse{}

	mockRepo.On("UpdateDelivery", request, id).Return(expectedResponse, nil)

	response, err := service.UpdateDelivery(request, id)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
	mockRepo.AssertExpectations(t)
}

func TestDeleteDelivery(t *testing.T) {
	mockRepo := new(MockDeliveryRepository)
	service := NewDeliveryService(mockRepo)

	id := 1

	mockRepo.On("DeleteDelivery", id).Return(nil)

	err := service.DeleteDelivery(id)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteAllDeliveries(t *testing.T) {
	mockRepo := new(MockDeliveryRepository)
	service := NewDeliveryService(mockRepo)

	mockRepo.On("DeleteAllDeliveries").Return(nil)

	err := service.DeleteAllDeliveries()

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
