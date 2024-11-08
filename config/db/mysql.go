package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Função que abre uma conexão com o banco de dados MySQL.
func OpenMySQLConnection() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
	}

	// Configurações do pool de conexões simultâneas
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Criando tabelas se não existirem
	err = InitTables(db)

	if err != nil {
		log.Fatalf("Erro ao criar tabelas: %v", err)
	}

	// Verifica se a conexão com o banco de dados está funcionando
	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao conectar-se ao banco de dados: %v", err)
	}

	return db
}

func InitTables(db *sql.DB) error {
	sql := `CREATE TABLE IF NOT EXISTS entregas (
    id INT PRIMARY KEY AUTO_INCREMENT,
    cliente VARCHAR(255) NOT NULL,
    peso FLOAT NOT NULL,
    endereco VARCHAR(255) NOT NULL,
    logradouro VARCHAR(255),
    numero VARCHAR(50),
    bairro VARCHAR(255),
    complemento VARCHAR(255),
    cidade VARCHAR(255) NOT NULL,
    estado VARCHAR(100) NOT NULL,
    pais VARCHAR(100) NOT NULL,
    latitude DOUBLE,
    longitude DOUBLE,
    data_inclusao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    data_alteracao TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP);`

	_, err := db.Exec(sql)

	return err
}
