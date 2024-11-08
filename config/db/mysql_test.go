package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// Teste de integração utilizando testcontainers para testar a conexão com o banco de dados

func setupMySQLContainer(t *testing.T) (*sql.DB, func()) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "password",
			"MYSQL_DATABASE":      "testdb",
		},
		WaitingFor: wait.ForSQL("3306/tcp", "mysql", func(host string, port nat.Port) string {
			return fmt.Sprintf("root:password@tcp(%s:%s)/testdb", host, port.Port())
		}).WithStartupTimeout(5 * time.Minute),
	}

	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Failed to start container: %v", err)
	}

	host, err := mysqlC.Host(ctx)
	if err != nil {
		t.Fatalf("Failed to get container host: %v", err)
	}

	port, err := mysqlC.MappedPort(ctx, "3306/tcp")
	if err != nil {
		t.Fatalf("Failed to get mapped port: %v", err)
	}

	dsn := fmt.Sprintf("root:password@tcp(%s:%s)/testdb", host, port.Port())
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}

	time.Sleep(5 * time.Second)

	err = db.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	teardown := func() {
		db.Close()
		mysqlC.Terminate(ctx)
	}

	return db, teardown
}

func TestInitTablesWithContainer(t *testing.T) {
	db, teardown := setupMySQLContainer(t)
	defer teardown()

	err := InitTables(db)
	assert.NoError(t, err)

	var tableExists bool
	row := db.QueryRow("SELECT COUNT(*) > 0 FROM information_schema.tables WHERE table_name = 'entregas'")
	err = row.Scan(&tableExists)

	assert.NoError(t, err)
	assert.True(t, tableExists)
}
