# É importante executar primeiro o docker composer

1 - No terminal executar o comando: docker compose up -d --build
#
2 - Em outra aba do terminal executa o camando: evans --host localhost --port 50051 -r repl
  - package pb
  - service OrderService
  - call ListOrders
#
3 - No browser acessar localhost:8080
    query ListOrders{
      orders {
        id
        price
        tax
        finalPrice
      }
    }

# Testes que você pode fazer:
- WEB SERVICE
  Executa o comando no api.http
- Listar
  GET http://localhost:8000/orders

- GRPC
  evans --host localhost --port 50051 -r repl
  package pb
  service OrderService
  call ListOrders

- GRAPHQL
  query ListOrders{
    orders {
      id
      price
      tax
      finalPrice
    }
  }

O projeto utiliza Docker para o ambiente de desenvolvimento, incluindo um banco de dados PostgreSQL.

## Tecnologias Utilizadas

# - Go
# - Mysql
# - Docker
# - gRPC 
  - https://grpc.io
  - https://developers.google.com/protocol-buffers
  - https://protobuf.dev/installation/
  1 - apt install -y protobuf-compiler - apos a instalacao executar o comando "protoc --version"
  2 - go get -u google.golang.org/protobuf/cmd/protoc-gen-go@latest
  3 - go get -u google.golang.org/protobuf/cmd/protoc-gen-go-grpc@latest

  # Comando para gerar as entidades no protocol buffer
  - protoc --go_out=. --go-grpc_out=. proto/order.proto

  # Evans
  - https://github.com/ktr0731/evans
  - go install github.com/ktr0731/evans@latest
  - evans -r repl

# - GraphQL (graphql-go) 
  - https://gqlgen.com/
  - printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
  - go mod tidy
  - go run github.com/99designs/gqlgen init
  - go mod tidy
  # apos a criação do schema
    - go run github.com/99designs/gqlgen generate

- REST (net/http ou gorilla/mux)
- SQLX para interação com o banco de dados
- Golang Migrate para migrações (opcionalmente integrado ao Docker)

## Pré-requisitos

- Docker e Docker Compose instalados.
- Go (v1.21 ou superior) instalado (para desenvolvimento local e geração de código gRPC).
- `protoc` (Protocol Buffer Compiler) e plugins `protoc-gen-go`, `protoc-gen-go-grpc`.
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  # Certifique-se de que $GOPATH/bin está no seu PATH

  go run main.go wire_gen.go