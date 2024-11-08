package delivery

import (
	"context"
	"database/sql"
)

type DeliveryRepository struct {
	db *sql.DB
}

type IDeliveryRepository interface {
	CreateDelivery(request *CreateDeliveryRequest) (*DeliveryResponse, error)
	UpdateDelivery(request *UpdateDeliveryRequest, id int) (*DeliveryResponse, error)
	GetDelivery(id int) (*DeliveryResponse, error)
	GetDeliveries() ([]*DeliveryResponse, error)
	GetDeliveriesByCity(city string) ([]*DeliveryResponse, error)
	DeleteDelivery(id int) error
	DeleteAllDeliveries() error
}

func NewDeliveryRepository(db *sql.DB) IDeliveryRepository {
	return &DeliveryRepository{db: db}
}

// Função responsável por inserir uma nova entrega no banco de dados.
func (r DeliveryRepository) CreateDelivery(request *CreateDeliveryRequest) (*DeliveryResponse, error) {
	// Criando contexto e transação para possibilitar rollback em caso de erro
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Executando a query usando o contexto e transação
	res, err := tx.ExecContext(ctx, insertDeliveryQuery,
		&request.Cliente,
		&request.Peso,
		&request.Endereco,
		&request.Logradouro,
		&request.Numero,
		&request.Bairro,
		&request.Complemento,
		&request.Cidade,
		&request.Estado,
		&request.Pais,
		&request.Latitude,
		&request.Longitude,
	)

	if err != nil {
		return nil, err
	}

	// Commitando a transação
	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	// Buscando o delivery recém-criado
	delivery, err := r.GetDelivery(int(id))

	if err != nil {
		return nil, err
	}

	return delivery, nil
}

// Função responsável por atualizar uma entrega pelo seu ID.
func (r DeliveryRepository) UpdateDelivery(request *UpdateDeliveryRequest, id int) (*DeliveryResponse, error) {
	// Criando contexto e transação para possibilitar rollback em caso de erro
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Executando a query usando o contexto e transação
	_, err = tx.ExecContext(ctx, updateDeliveryQuery,
		&request.Peso,
		&request.Endereco,
		&request.Logradouro,
		&request.Numero,
		&request.Bairro,
		&request.Complemento,
		&request.Cidade,
		&request.Estado,
		&request.Pais,
		&request.Latitude,
		&request.Longitude,
		id,
	)

	if err != nil {
		// Verificando se o erro aconteceu por não encontrar a entrega para atualização
		if err == sql.ErrNoRows {
			return nil, ErrDeliveryNotFound
		}
		return nil, err
	}

	// Commitando a transação
	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	// Buscando o delivery recém-atualizado
	delivery, err := r.GetDelivery(int(id))

	if err != nil {
		return nil, err
	}

	return delivery, nil
}

// Função responsável por buscar uma entrega pelo seu ID.
func (r DeliveryRepository) GetDelivery(id int) (*DeliveryResponse, error) {
	var delivery Delivery

	// Executando a query de consulta sem necessidade de transação
	err := r.db.QueryRow(getDeliveryQuery, id).Scan(
		&delivery.ID,
		&delivery.Cliente,
		&delivery.Peso,
		&delivery.Endereco,
		&delivery.Logradouro,
		&delivery.Numero,
		&delivery.Bairro,
		&delivery.Complemento,
		&delivery.Cidade,
		&delivery.Estado,
		&delivery.Pais,
		&delivery.Latitude,
		&delivery.Longitude,
		&delivery.DataInclusao,
		&delivery.DataAlteracao,
	)

	if err != nil {
		// Verificando se o erro aconteceu por não encontrar a entrega
		if err == sql.ErrNoRows {
			return nil, ErrDeliveryNotFound
		}
		return nil, err
	}

	// Convertendo o model para o response
	response := delivery.ToDeliveryResponse()

	return response, nil
}

// Função responsável por buscar todas as entregas.
func (r DeliveryRepository) GetDeliveries() ([]*DeliveryResponse, error) {
	var deliveries []*DeliveryResponse = make([]*DeliveryResponse, 0)

	// Executando a query de consulta sem necessidade de transação
	rows, err := r.db.Query(getDeliveriesQuery)

	if err != nil {
		return nil, err
	}

	// Fechando a conexão com o cursor em caso de erro
	defer rows.Close()

	// Iterando sobre os resultados da consulta
	for rows.Next() {
		var delivery Delivery

		// Escaneando os resultados da consulta para o model
		err := rows.Scan(
			&delivery.ID,
			&delivery.Cliente,
			&delivery.Peso,
			&delivery.Endereco,
			&delivery.Logradouro,
			&delivery.Numero,
			&delivery.Bairro,
			&delivery.Complemento,
			&delivery.Cidade,
			&delivery.Estado,
			&delivery.Pais,
			&delivery.Latitude,
			&delivery.Longitude,
			&delivery.DataInclusao,
			&delivery.DataAlteracao,
		)

		if err != nil {
			return nil, err
		}

		// Convertendo o model para o response e adicionando ao array
		deliveries = append(deliveries, delivery.ToDeliveryResponse())
	}

	return deliveries, nil
}

// Função responsável por buscar todas as entregas filtrando por cidade.
func (r DeliveryRepository) GetDeliveriesByCity(city string) ([]*DeliveryResponse, error) {
	var deliveries []*DeliveryResponse = make([]*DeliveryResponse, 0)

	// Executando a query de consulta sem necessidade de transação
	rows, err := r.db.Query(getDeliveriesByCityQuery, city)

	if err != nil {
		return nil, err
	}

	// Fechando a conexão com o cursor em caso de erro
	defer rows.Close()

	// Iterando sobre os resultados da consulta
	for rows.Next() {
		var delivery Delivery

		// Escaneando os resultados da consulta para o model
		err := rows.Scan(
			&delivery.ID,
			&delivery.Cliente,
			&delivery.Peso,
			&delivery.Endereco,
			&delivery.Logradouro,
			&delivery.Numero,
			&delivery.Bairro,
			&delivery.Complemento,
			&delivery.Cidade,
			&delivery.Estado,
			&delivery.Pais,
			&delivery.Latitude,
			&delivery.Longitude,
			&delivery.DataInclusao,
			&delivery.DataAlteracao,
		)

		if err != nil {
			return nil, err
		}

		// Convertendo o model para o response e adicionando ao array
		deliveries = append(deliveries, delivery.ToDeliveryResponse())
	}

	return deliveries, nil
}

// Função responsável por excluir uma entrega pelo seu ID.
func (r DeliveryRepository) DeleteDelivery(id int) error {
	// Criando contexto e transação para possibilitar rollback em caso de erro
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	// Executando a query usando o contexto e transação
	_, err = tx.ExecContext(ctx, deleteDeliveryQuery, id)

	if err != nil {
		// Verificando se o erro aconteceu por não encontrar a entrega para exclusão
		if err == sql.ErrNoRows {
			return ErrDeliveryNotFound
		}
		return err
	}

	// Commitando a transação
	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

// Função responsável por excluir todas as entregas.
func (r DeliveryRepository) DeleteAllDeliveries() error {
	// Criando contexto e transação para possibilitar rollback em caso de erro
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	// Executando a query usando o contexto e transação
	_, err = tx.ExecContext(ctx, deleteAllDeliveriesQuery)

	if err != nil {
		return err
	}

	// Commitando a transação
	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}
