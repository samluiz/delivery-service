package delivery

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(*sql.Tx), args.Error(1)
}

func (m *MockDB) QueryRow(query string, args ...interface{}) *sql.Row {
	call := m.Called(query, args)
	return call.Get(0).(*sql.Row)
}

func (m *MockDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	call := m.Called(ctx, query, args)
	return call.Get(0).(sql.Result), call.Error(1)
}

func TestCreateDeliveryRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewDeliveryRepository(db)

	request := &CreateDeliveryRequest{
		Cliente:     "Cliente A",
		Peso:        10.5,
		Endereco:    "Endereço 123",
		Logradouro:  "Rua 1",
		Numero:      "123",
		Bairro:      "Bairro A",
		Complemento: "Casa",
		Cidade:      "Cidade A",
		Estado:      "Estado A",
		Pais:        "País A",
		Latitude:    40.7128,
		Longitude:   -74.0060,
	}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO entregas`).
		WithArgs(request.Cliente, request.Peso, request.Endereco, request.Logradouro, request.Numero, request.Bairro, request.Complemento, request.Cidade, request.Estado, request.Pais, request.Latitude, request.Longitude).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectQuery(`SELECT \* FROM entregas WHERE id = \?`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "cliente", "peso", "endereco", "logradouro", "numero", "bairro", "complemento", "cidade", "estado", "pais", "latitude", "longitude", "data_inclusao", "data_alteracao"}).
			AddRow(1, "Cliente A", 10.5, "Endereço 123", "Rua 1", "123", "Bairro A", "Casa", "Cidade A", "Estado A", "País A", 40.7128, -74.0060, time.Now(), time.Now()))

	delivery, err := repo.CreateDelivery(request)
	assert.NoError(t, err)
	assert.Equal(t, "Cliente A", delivery.Cliente)
}

func TestUpdateDeliveryRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewDeliveryRepository(db)

	request := &UpdateDeliveryRequest{
		Peso:        12.5,
		Endereco:    "456 Novo Endereço",
		Logradouro:  "Nova Rua",
		Numero:      "456",
		Bairro:      "Novo Bairro",
		Complemento: "Apartamento",
		Cidade:      "Nova Cidade",
		Estado:      "Novo Estado",
		Pais:        "Novo País",
		Latitude:    51.5074,
		Longitude:   -0.1278,
	}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE entregas`).
		WithArgs(request.Peso, request.Endereco, request.Logradouro, request.Numero, request.Bairro, request.Complemento, request.Cidade, request.Estado, request.Pais, request.Latitude, request.Longitude, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectQuery(`SELECT \* FROM entregas WHERE id = \?`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "cliente", "peso", "endereco", "logradouro", "numero", "bairro", "complemento", "cidade", "estado", "pais", "latitude", "longitude", "data_inclusao", "data_alteracao"}).
			AddRow(1, "Cliente A", 12.5, "456 Novo Endereço", "Nova Rua", "456", "Novo Bairro", "Apartamento", "Nova Cidade", "Novo Estado", "Novo País", 51.5074, -0.1278, time.Now(), time.Now()))

	delivery, err := repo.UpdateDelivery(request, 1)
	assert.NoError(t, err)
	assert.Equal(t, "Nova Cidade", delivery.Cidade)
}

func TestGetDeliveryRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewDeliveryRepository(db)

	mock.ExpectQuery(`SELECT \* FROM entregas WHERE id = \?`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "cliente", "peso", "endereco", "logradouro", "numero", "bairro", "complemento", "cidade", "estado", "pais", "latitude", "longitude", "data_inclusao", "data_alteracao"}).
			AddRow(1, "Cliente A", 10.5, "Endereço 123", "Rua 1", "123", "Bairro A", "Casa", "Cidade A", "Estado A", "País A", 40.7128, -74.0060, time.Now(), time.Now()))

	delivery, err := repo.GetDelivery(1)
	assert.NoError(t, err)
	assert.Equal(t, "Cliente A", delivery.Cliente)
}

func TestGetDeliveriesRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewDeliveryRepository(db)

	mock.ExpectQuery(`SELECT \* FROM entregas ORDER BY id DESC`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "cliente", "peso", "endereco", "logradouro", "numero", "bairro", "complemento", "cidade", "estado", "pais", "latitude", "longitude", "data_inclusao", "data_alteracao"}).
			AddRow(1, "Cliente A", 10.5, "Endereço 123", "Rua 1", "123", "Bairro A", "Casa", "Cidade A", "Estado A", "País A", 40.7128, -74.0060, time.Now(), time.Now()).
			AddRow(2, "Cliente B", 20.0, "Endereço 456", "Rua 2", "456", "Bairro B", "Apartamento", "Cidade B", "Estado B", "País B", 51.5074, -0.1278, time.Now(), time.Now()))

	deliveries, err := repo.GetDeliveries()
	assert.NoError(t, err)
	assert.Len(t, deliveries, 2)
}

func TestGetDeliveriesByCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewDeliveryRepository(db)

	mock.ExpectQuery(`SELECT \* FROM entregas WHERE cidade = \? ORDER BY id DESC`).
		WithArgs("São Paulo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "cliente", "peso", "endereco", "logradouro", "numero", "bairro", "complemento", "cidade", "estado", "pais", "latitude", "longitude", "data_inclusao", "data_alteracao"}).
			AddRow(2, "Cliente B", 20.0, "Endereço 456", "Rua 2", "456", "Bairro B", "Apartamento", "São Paulo", "Estado B", "País B", 51.5074, -0.1278, time.Now(), time.Now()))

	deliveries, err := repo.GetDeliveriesByCity("São Paulo")
	assert.NoError(t, err)
	assert.Len(t, deliveries, 1)
}

func TestDeleteDeliveryRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewDeliveryRepository(db)

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM entregas WHERE id = \?`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = repo.DeleteDelivery(1)
	assert.NoError(t, err)
}

func TestDeleteAllDeliveriesRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewDeliveryRepository(db)

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM entregas`).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = repo.DeleteAllDeliveries()
	assert.NoError(t, err)
}

func TestCreateDelivery_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin().WillReturnError(errors.New("transaction error"))

	repo := NewDeliveryRepository(db)

	request := &CreateDeliveryRequest{}
	delivery, err := repo.CreateDelivery(request)

	assert.Nil(t, delivery)
	assert.Error(t, err)
	assert.Equal(t, "transaction error", err.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateDelivery_ExecError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO entregas (
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
		)`)).WillReturnError(errors.New("exec error"))

	repo := NewDeliveryRepository(db)

	request := &CreateDeliveryRequest{}
	delivery, err := repo.CreateDelivery(request)

	assert.Nil(t, delivery)
	assert.Error(t, err)
	assert.Equal(t, "exec error", err.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateDelivery_CommitError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO entregas (
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
		)`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit().WillReturnError(errors.New("commit error"))

	repo := NewDeliveryRepository(db)

	request := &CreateDeliveryRequest{}
	delivery, err := repo.CreateDelivery(request)

	assert.Nil(t, delivery)
	assert.Error(t, err)
	assert.Equal(t, "commit error", err.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateDelivery_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin().WillReturnError(errors.New("transaction error"))

	repo := NewDeliveryRepository(db)

	request := &UpdateDeliveryRequest{}
	delivery, err := repo.UpdateDelivery(request, 1)

	assert.Nil(t, delivery)
	assert.Error(t, err)
	assert.Equal(t, "transaction error", err.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateDelivery_ExecError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE entregas`).WillReturnError(errors.New("exec error"))

	repo := NewDeliveryRepository(db)

	request := &UpdateDeliveryRequest{}
	delivery, err := repo.UpdateDelivery(request, 1)

	assert.Nil(t, delivery)
	assert.Error(t, err)
	assert.Equal(t, "exec error", err.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateDelivery_CommitError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE entregas SET 
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
		WHERE id = ?`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit().WillReturnError(errors.New("commit error"))

	repo := NewDeliveryRepository(db)

	request := &UpdateDeliveryRequest{}
	delivery, err := repo.UpdateDelivery(request, 1)

	assert.Nil(t, delivery)
	assert.Error(t, err)
	assert.Equal(t, "commit error", err.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteDelivery_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin().WillReturnError(errors.New("transaction error"))

	repo := NewDeliveryRepository(db)

	err = repo.DeleteDelivery(1)

	assert.Error(t, err)
	assert.Equal(t, "transaction error", err.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteDelivery_ExecError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM entregas WHERE id = ?`)).WillReturnError(errors.New("exec error"))

	repo := NewDeliveryRepository(db)

	err = repo.DeleteDelivery(1)

	assert.Error(t, err)
	assert.Equal(t, "exec error", err.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteDelivery_CommitError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM entregas WHERE id = ?`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit().WillReturnError(errors.New("commit error"))

	repo := NewDeliveryRepository(db)

	err = repo.DeleteDelivery(1)

	assert.Error(t, err)
	assert.Equal(t, "commit error", err.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetDelivery_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM entregas WHERE id = ?`)).WillReturnError(errors.New("query error"))

	repo := NewDeliveryRepository(db)

	delivery, err := repo.GetDelivery(1)

	assert.Nil(t, delivery)
	assert.Error(t, err)
	assert.Equal(t, "query error", err.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetDeliveries_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM entregas`)).WillReturnError(errors.New("query error"))

	repo := NewDeliveryRepository(db)

	deliveries, err := repo.GetDeliveries()

	assert.Nil(t, deliveries)
	assert.Error(t, err)
	assert.Equal(t, "query error", err.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteAllDeliveries_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin().WillReturnError(errors.New("transaction error"))

	repo := NewDeliveryRepository(db)

	err = repo.DeleteAllDeliveries()

	assert.Error(t, err)
	assert.Equal(t, "transaction error", err.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteAllDeliveries_ExecError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM entregas`)).WillReturnError(errors.New("exec error"))

	repo := NewDeliveryRepository(db)

	err = repo.DeleteAllDeliveries()

	assert.Error(t, err)
	assert.Equal(t, "exec error", err.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteAllDeliveries_CommitError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM entregas`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit().WillReturnError(errors.New("commit error"))

	repo := NewDeliveryRepository(db)

	err = repo.DeleteAllDeliveries()

	assert.Error(t, err)
	assert.Equal(t, "commit error", err.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}
