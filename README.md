# Delivery Service API

API de gerenciamento de entregas com geolocalização


## Funcionalidades

- Criação, atualização, visualização e remoção de entregas
- Documentação Swagger
- Testes unitários


## Rodando localmente

Clone o projeto

```bash
  git clone https://github.com/samluiz/delivery-service
```

Entre no diretório do projeto

```bash
  cd delivery-service
```

Inicie o servidor (é necessário possuir o Docker e docker-compose instalados em sua máquina)

```bash
  docker-compose up
```


## Rodando os testes

Para rodar os testes, rode o seguinte comando

```bash
  go test ./... -cover
```


## Stack

- Go (1.23.0)
- MySQL

#### Bibliotecas
- github.com/go-playground/validator/v10 (para validar as structs das requisições)
- github.com/go-sql-driver/mysql (driver do banco de dados)
- github.com/stretchr/testify (para os testes unitários)
- github.com/DATA-DOG/go-sqlmock (para mockar os testes unitários do repositório)
- github.com/docker/go-connections (para utilizar testcontainers nos testes de integração)
- github.com/testcontainers/testcontainers-go

#### Ferramentas
- Docker
- Docker Compose