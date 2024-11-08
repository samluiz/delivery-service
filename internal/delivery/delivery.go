package delivery

import "time"

var (
	TABLE_NAME = "entregas"
)

type Delivery struct {
	ID            int       `db:"id"`
	Cliente       string    `db:"cliente"`
	Peso          float64   `db:"peso"`
	Endereco      string    `db:"endereco"`
	Logradouro    string    `db:"logradouro"`
	Numero        string    `db:"numero"`
	Bairro        string    `db:"bairro"`
	Complemento   string    `db:"complemento"`
	Cidade        string    `db:"cidade"`
	Estado        string    `db:"estado"`
	Pais          string    `db:"pais"`
	Latitude      float64   `db:"latitude"`
	Longitude     float64   `db:"longitude"`
	DataInclusao  time.Time `db:"data_inclusao"`
	DataAlteracao time.Time `db:"data_alteracao"`
}

type CreateDeliveryRequest struct {
	Cliente     string  `json:"cliente" validate:"required"`
	Peso        float64 `json:"peso" validate:"required"`
	Endereco    string  `json:"endereco" validate:"required"`
	Logradouro  string  `json:"logradouro" validate:"required"`
	Numero      string  `json:"numero" validate:"required"`
	Bairro      string  `json:"bairro" validate:"required"`
	Complemento string  `json:"complemento" validate:"required"`
	Cidade      string  `json:"cidade" validate:"required"`
	Estado      string  `json:"estado" validate:"required"`
	Pais        string  `json:"pais" validate:"required"`
	Latitude    float64 `json:"latitude" validate:"required"`
	Longitude   float64 `json:"longitude" validate:"required"`
}

type UpdateDeliveryRequest struct {
	Peso        float64 `json:"peso" validate:"required"`
	Endereco    string  `json:"endereco" validate:"required"`
	Logradouro  string  `json:"logradouro" validate:"required"`
	Numero      string  `json:"numero" validate:"required"`
	Bairro      string  `json:"bairro" validate:"required"`
	Complemento string  `json:"complemento" validate:"required"`
	Cidade      string  `json:"cidade" validate:"required"`
	Estado      string  `json:"estado" validate:"required"`
	Pais        string  `json:"pais" validate:"required"`
	Latitude    float64 `json:"latitude" validate:"required"`
	Longitude   float64 `json:"longitude" validate:"required"`
}

type DeliveryResponse struct {
	ID            int       `json:"id"`
	Cliente       string    `json:"cliente"`
	Peso          float64   `json:"peso"`
	Endereco      string    `json:"endereco"`
	Logradouro    string    `json:"logradouro"`
	Numero        string    `json:"numero"`
	Bairro        string    `json:"bairro"`
	Complemento   string    `json:"complemento"`
	Cidade        string    `json:"cidade"`
	Estado        string    `json:"estado"`
	Pais          string    `json:"pais"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	DataInclusao  time.Time `json:"data_inclusao"`
	DataAlteracao time.Time `json:"data_alteracao"`
}

func (r DeliveryResponse) ToDelivery() *Delivery {
	return &Delivery{
		ID:            r.ID,
		Cliente:       r.Cliente,
		Peso:          r.Peso,
		Endereco:      r.Endereco,
		Logradouro:    r.Logradouro,
		Numero:        r.Numero,
		Bairro:        r.Bairro,
		Complemento:   r.Complemento,
		Cidade:        r.Cidade,
		Estado:        r.Estado,
		Pais:          r.Pais,
		Latitude:      r.Latitude,
		Longitude:     r.Longitude,
		DataInclusao:  r.DataInclusao,
		DataAlteracao: r.DataAlteracao,
	}
}

func (r Delivery) ToDeliveryResponse() *DeliveryResponse {
	return &DeliveryResponse{
		ID:            r.ID,
		Cliente:       r.Cliente,
		Peso:          r.Peso,
		Endereco:      r.Endereco,
		Logradouro:    r.Logradouro,
		Numero:        r.Numero,
		Bairro:        r.Bairro,
		Complemento:   r.Complemento,
		Cidade:        r.Cidade,
		Estado:        r.Estado,
		Pais:          r.Pais,
		Latitude:      r.Latitude,
		Longitude:     r.Longitude,
		DataInclusao:  r.DataInclusao,
		DataAlteracao: r.DataAlteracao,
	}
}

var (
	insertDeliveryQuery = `INSERT INTO entregas (
			cliente,
			peso,
			endereco,
			logradouro,
			numero,
			bairro,
			complemento,
			cidade,
			estado,
			pais,
			latitude,
			longitude
		) VALUES (
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?
		)`

	updateDeliveryQuery = `UPDATE entregas
			SET
				cliente = ?,
				peso = ?,
				endereco = ?,
				logradouro = ?,
				numero = ?,
				bairro = ?,
				complemento = ?,
				cidade = ?,
				estado = ?,
				pais = ?,
				latitude = ?,
				longitude = ?
			WHERE id = ?`

	getDeliveryQuery = `SELECT * FROM entregas WHERE id = ?`

	getDeliveriesQuery = `SELECT * FROM entregas ORDER BY id DESC`

	getDeliveriesByCityQuery = `SELECT * FROM entregas WHERE cidade = ? ORDER BY id DESC`

	deleteDeliveryQuery = `DELETE FROM entregas WHERE id = ?`

	deleteAllDeliveriesQuery = `DELETE FROM entregas`
)
