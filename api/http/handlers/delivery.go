package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/samluiz/delivery-service/api/http/utils"
	"github.com/samluiz/delivery-service/internal/delivery"
)

type DeliveryHandler struct {
	deliveryService delivery.IDeliveryService
}

func NewDeliveryHandler(deliveryService delivery.IDeliveryService) *DeliveryHandler {
	return &DeliveryHandler{deliveryService: deliveryService}
}

func (h DeliveryHandler) HandleCreateDelivery(w http.ResponseWriter, r *http.Request) {
	var request delivery.CreateDeliveryRequest

	// Serializando o request body para o struct
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.NewJSONResponse(w, http.StatusBadRequest, utils.NewBadRequestError(err, r))
		return
	}

	// Validando o request body
	validationError := utils.ValidateBody(r, &request)

	if validationError != nil {
		utils.NewJSONResponse(w, http.StatusBadRequest, validationError)
		return
	}

	response, err := h.deliveryService.CreateDelivery(&request)

	if err != nil {
		utils.NewJSONResponse(w, http.StatusInternalServerError, utils.NewInternalServerError(err, r))
		return
	}

	utils.NewJSONResponse(w, http.StatusCreated, response)
}

func (h DeliveryHandler) HandleGetDelivery(w http.ResponseWriter, r *http.Request) {
	// Buscando o ID da entrega no path
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		utils.NewJSONResponse(w, http.StatusBadRequest, utils.NewBadRequestError(err, r))
		return
	}

	response, err := h.deliveryService.GetDelivery(id)

	if err != nil {
		// Verificando se o erro aconteceu por não encontrar a entrega
		if errors.Is(err, delivery.ErrDeliveryNotFound) {
			// Retornando o erro de não encontrado
			utils.NewJSONResponse(w, http.StatusNotFound, utils.NewNotFoundError(err, r))
			return
		}
		utils.NewJSONResponse(w, http.StatusInternalServerError, utils.NewInternalServerError(err, r))
		return
	}

	utils.NewJSONResponse(w, http.StatusOK, response)
}

func (h DeliveryHandler) HandleGetDeliveries(w http.ResponseWriter, r *http.Request) {
	// Buscando o query param de cidade
	city := r.URL.Query().Get("city")

	response, err := h.deliveryService.GetDeliveries(city)

	if err != nil {
		utils.NewJSONResponse(w, http.StatusInternalServerError, utils.NewInternalServerError(err, r))
		return
	}

	utils.NewJSONResponse(w, http.StatusOK, response)
}

func (h DeliveryHandler) HandleUpdateDelivery(w http.ResponseWriter, r *http.Request) {
	// Buscando o ID da entrega no path
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		utils.NewJSONResponse(w, http.StatusBadRequest, utils.NewBadRequestError(err, r))
		return
	}

	var request delivery.UpdateDeliveryRequest

	// Serializando o request body para o struct
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.NewJSONResponse(w, http.StatusBadRequest, utils.NewBadRequestError(err, r))
		return
	}

	// Validando o request body
	validationError := utils.ValidateBody(r, &request)

	if validationError != nil {
		utils.NewJSONResponse(w, http.StatusBadRequest, validationError)
		return
	}

	response, err := h.deliveryService.UpdateDelivery(&request, id)

	if err != nil {
		// Verificando se o erro aconteceu por não encontrar a entrega
		if errors.Is(err, delivery.ErrDeliveryNotFound) {
			// Retornando o erro de não encontrado
			utils.NewJSONResponse(w, http.StatusNotFound, utils.NewNotFoundError(err, r))
			return
		}
		utils.NewJSONResponse(w, http.StatusInternalServerError, utils.NewInternalServerError(err, r))
		return
	}

	utils.NewJSONResponse(w, http.StatusOK, response)
}

func (h DeliveryHandler) HandleDeleteDelivery(w http.ResponseWriter, r *http.Request) {
	// Buscando o ID da entrega no path
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		utils.NewJSONResponse(w, http.StatusBadRequest, utils.NewBadRequestError(err, r))
		return
	}

	err = h.deliveryService.DeleteDelivery(id)

	if err != nil {
		// Verificando se o erro aconteceu por não encontrar a entrega
		if errors.Is(err, delivery.ErrDeliveryNotFound) {
			// Retornando o erro de não encontrado
			utils.NewJSONResponse(w, http.StatusNotFound, utils.NewNotFoundError(err, r))
			return
		}
		utils.NewJSONResponse(w, http.StatusInternalServerError, utils.NewInternalServerError(err, r))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h DeliveryHandler) HandleDeleteAllDeliveries(w http.ResponseWriter, r *http.Request) {
	err := h.deliveryService.DeleteAllDeliveries()

	if err != nil {
		utils.NewJSONResponse(w, http.StatusInternalServerError, utils.NewInternalServerError(err, r))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
