
# Go Book API

![Go](https://img.shields.io/badge/Go-100%25-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Build](https://img.shields.io/github/actions/workflow/status/mateushenriquedasilva/go-book-api/.github/workflows/go.yml)

API simples para gerenciamento de livros, desenvolvida em Go.

---

## Funcionalidades

- Cadastro, listagem, atualização e remoção de livros
- Persistência de dados em banco relacional
- Estrutura modular (handlers, models, database)

---

## Instalação

1. Clone o repositório:
	 ```bash
	 git clone https://github.com/mateushenriquedasilva/go-book-api.git
	 cd go-book-api
	 ```
2. Instale as dependências:
	 ```bash
	 go mod tidy
	 ```
3. Configure o banco de dados conforme necessário.

---

## Como Usar

Para iniciar a API:
```bash
go run ./cmd/main.go
```

---

## Exemplos de Endpoints

### Listar livros
```http
GET /books
```

### Adicionar livro
```http
POST /books
Content-Type: application/json
{
	"title": "Nome do Livro",
	"author": "Autor"
}
```

### Atualizar livro
```http
PUT /books/{id}
Content-Type: application/json
{
	"title": "Novo Título",
	"author": "Novo Autor"
}
```

### Remover livro
```http
DELETE /books/{id}
```

---

## Estrutura do Projeto

```
go-book-api/
├── api/
│   ├── database.go
│   ├── handlers.go
│   └── model.go
├── cmd/
│   └── main.go
├── go.mod
├── go.sum
└── README.md
```

---

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).

---

## Contato

Para dúvidas ou sugestões, entre em contato:

- [matheus.hsilvaa18@gmail.com](mailto:matheus.hsilvaa18@gmail.com)
- [Matheus Henrique da Silva](https://github.com/mateushenriquedasilva)
