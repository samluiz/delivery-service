package delivery

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Testes de campos do request

func TestCreateDeliveryRequest(t *testing.T) {
	request := &CreateDeliveryRequest{
		Cliente:     "Cliente 1",
		Peso:        10.5,
		Endereco:    "Endereço 123",
		Logradouro:  "Rua 1",
		Numero:      "123",
		Bairro:      "Bairro 2",
		Complemento: "Apartamento 3",
		Cidade:      "Cidade 4",
		Estado:      "Estado 5",
		Pais:        "País 6",
		Latitude:    40.7128,
		Longitude:   -74.0060,
	}

	assert.Equal(t, request.Cliente, "Cliente 1")
	assert.Equal(t, request.Peso, 10.5)
	assert.Equal(t, request.Endereco, "Endereço 123")
	assert.Equal(t, request.Logradouro, "Rua 1")
	assert.Equal(t, request.Numero, "123")
	assert.Equal(t, request.Bairro, "Bairro 2")
	assert.Equal(t, request.Complemento, "Apartamento 3")
	assert.Equal(t, request.Cidade, "Cidade 4")
	assert.Equal(t, request.Estado, "Estado 5")
	assert.Equal(t, request.Pais, "País 6")
	assert.Equal(t, request.Latitude, 40.7128)
	assert.Equal(t, request.Longitude, -74.0060)
}

func TestUpdateDeliveryRequest(t *testing.T) {
	request := &UpdateDeliveryRequest{
		Peso:        12.5,
		Endereco:    "456 Novo Endereço",
		Logradouro:  "Nova Rua",
		Numero:      "456",
		Bairro:      "Novo Bairro",
		Complemento: "Casa",
		Cidade:      "Nova Cidade",
		Estado:      "Novo Estado",
		Pais:        "Novo País",
		Latitude:    51.5074,
		Longitude:   -0.1278,
	}

	assert.Equal(t, request.Peso, 12.5)
	assert.Equal(t, request.Endereco, "456 Novo Endereço")
	assert.Equal(t, request.Logradouro, "Nova Rua")
	assert.Equal(t, request.Numero, "456")
	assert.Equal(t, request.Bairro, "Novo Bairro")
	assert.Equal(t, request.Complemento, "Casa")
	assert.Equal(t, request.Cidade, "Nova Cidade")
	assert.Equal(t, request.Estado, "Novo Estado")
	assert.Equal(t, request.Pais, "Novo País")
	assert.Equal(t, request.Latitude, 51.5074)
	assert.Equal(t, request.Longitude, -0.1278)
}

// Testes de conversão de structs (model -> response)

func TestToDelivery(t *testing.T) {
	now := time.Now()

	request := &DeliveryResponse{
		ID:            1,
		Cliente:       "Cliente 1",
		Peso:          10.5,
		Endereco:      "Endereço 123",
		Logradouro:    "Rua 1",
		Numero:        "123",
		Bairro:        "Bairro 2",
		Complemento:   "Apartamento 3",
		Cidade:        "Cidade 4",
		Estado:        "Estado 5",
		Pais:          "País 6",
		Latitude:      40.7128,
		Longitude:     -74.0060,
		DataInclusao:  now,
		DataAlteracao: now,
	}

	delivery := request.ToDelivery()

	assert.Equal(t, delivery.ID, 1)
	assert.Equal(t, delivery.Cliente, "Cliente 1")
	assert.Equal(t, delivery.Peso, 10.5)
	assert.Equal(t, delivery.Endereco, "Endereço 123")
	assert.Equal(t, delivery.Logradouro, "Rua 1")
	assert.Equal(t, delivery.Numero, "123")
	assert.Equal(t, delivery.Bairro, "Bairro 2")
	assert.Equal(t, delivery.Complemento, "Apartamento 3")
	assert.Equal(t, delivery.Cidade, "Cidade 4")
	assert.Equal(t, delivery.Estado, "Estado 5")
	assert.Equal(t, delivery.Pais, "País 6")
	assert.Equal(t, delivery.Latitude, 40.7128)
	assert.Equal(t, delivery.Longitude, -74.0060)
	assert.Equal(t, delivery.DataInclusao, now)
	assert.Equal(t, delivery.DataAlteracao, now)
}

func TestToDeliveryResponse(t *testing.T) {
	now := time.Now()

	delivery := &Delivery{
		ID:            1,
		Cliente:       "Cliente 1",
		Peso:          10.5,
		Endereco:      "Endereço 123",
		Logradouro:    "Rua 1",
		Numero:        "123",
		Bairro:        "Bairro 2",
		Complemento:   "Apartamento 3",
		Cidade:        "Cidade 4",
		Estado:        "Estado 5",
		Pais:          "País 6",
		Latitude:      40.7128,
		Longitude:     -74.0060,
		DataInclusao:  now,
		DataAlteracao: now,
	}

	response := delivery.ToDeliveryResponse()

	assert.Equal(t, response.ID, 1)
	assert.Equal(t, response.Cliente, "Cliente 1")
	assert.Equal(t, response.Peso, 10.5)
	assert.Equal(t, response.Endereco, "Endereço 123")
	assert.Equal(t, response.Logradouro, "Rua 1")
	assert.Equal(t, response.Numero, "123")
	assert.Equal(t, response.Bairro, "Bairro 2")
	assert.Equal(t, response.Complemento, "Apartamento 3")
	assert.Equal(t, response.Cidade, "Cidade 4")
	assert.Equal(t, response.Estado, "Estado 5")
	assert.Equal(t, response.Pais, "País 6")
	assert.Equal(t, response.Latitude, 40.7128)
	assert.Equal(t, response.Longitude, -74.0060)
	assert.Equal(t, response.DataInclusao, now)
	assert.Equal(t, response.DataAlteracao, now)
}
