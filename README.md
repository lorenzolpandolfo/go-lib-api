# Go Library API

Uma API simples de cadastro e consulta de livros.

## Setup

Para iniciar o banco de dados, execute:

```bash
docker compose up -d
```

E então, para iniciar a API:

```bash
go mod tidy
go run cmd/api/main.go
```

---

# Exemplos de chamadas

Registrando um livro:

```bash
curl --location 'localhost:8080/books' \
--header 'Content-Type: application/json' \
--data '{
    "title": "1984",
    "author": "George Orwell"
}'
```

Existe uma regra de negócio que não permite um author chamado John

---

Consultando um livro por id

```bash
curl --location 'localhost:8080/books/1'
```

---
