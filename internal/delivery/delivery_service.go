package delivery

type DeliveryService struct {
	repository IDeliveryRepository
}

type IDeliveryService interface {
	CreateDelivery(request *CreateDeliveryRequest) (*DeliveryResponse, error)
	GetDelivery(id int) (*DeliveryResponse, error)
	GetDeliveries(city string) ([]*DeliveryResponse, error)
	UpdateDelivery(request *UpdateDeliveryRequest, id int) (*DeliveryResponse, error)
	DeleteDelivery(id int) error
	DeleteAllDeliveries() error
}

func NewDeliveryService(repository IDeliveryRepository) IDeliveryService {
	return &DeliveryService{repository: repository}
}

func (s DeliveryService) CreateDelivery(request *CreateDeliveryRequest) (*DeliveryResponse, error) {
	return s.repository.CreateDelivery(request)
}

func (s DeliveryService) GetDelivery(id int) (*DeliveryResponse, error) {
	return s.repository.GetDelivery(id)
}

func (s DeliveryService) GetDeliveries(city string) ([]*DeliveryResponse, error) {
	if city == "" {
		return s.repository.GetDeliveries()
	}
	return s.repository.GetDeliveriesByCity(city)
}

func (s DeliveryService) UpdateDelivery(request *UpdateDeliveryRequest, id int) (*DeliveryResponse, error) {
	return s.repository.UpdateDelivery(request, id)
}

func (s DeliveryService) DeleteDelivery(id int) error {
	return s.repository.DeleteDelivery(id)
}

func (s DeliveryService) DeleteAllDeliveries() error {
	return s.repository.DeleteAllDeliveries()
}
