# Estrutura do Projeto

```text
rupi/
│
├── cmd/
│   └── api/
│       └── main.go          # Ponto de entrada da aplicação
│
├── internal/
│   ├── config/
│   │   └── config.go
│   │
│   ├── database/
│   │   ├── postgres.go
│   │   └── migrations/
│   │
│   ├── urls/
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   ├── model.go
│   │   └── dto.go
│   │
│   ├── <other_domain>/
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── model.go
│   │
│   └── middleware/
│       ├── auth.go
│       └── logger.go
│
├── pkg/
│   ├── hashid/
│   └── validator/
│
├── scripts/
│
├── .env
├── go.mod
├── go.sum
└── Makefile
```

## Diretórios

### `cmd/`

Contém os executáveis da aplicação. Cada subdiretório representa um binário.

```text
cmd/
└── api/
    └── main.go
```

O `main.go` é responsável apenas por inicializar a aplicação:

- Carregar configurações;
- Criar conexões (banco, cache, etc.);
- Registrar dependências;
- Inicializar o servidor HTTP.

---

### `internal/`

Contém toda a lógica interna da aplicação.

Pacotes dentro de `internal` não podem ser importados por projetos externos.

---

### `internal/config/`

Responsável por carregar as configurações da aplicação.

Exemplo:

- Variáveis de ambiente;
- Porta da aplicação;
- URL do banco;
- Chaves secretas.

---

### `internal/database/`

Responsável pela conexão com o banco de dados.

Pode conter:

- criação do pool de conexões;
- migrations;
- helpers relacionados ao banco.

---

### `internal/shorturl/`

Domínio responsável pelo encurtamento de URLs.

Arquivos:

| Arquivo | Responsabilidade |
|----------|------------------|
| `handler.go` | Recebe requisições HTTP |
| `service.go` | Regras de negócio |
| `repository.go` | Comunicação com o banco |
| `model.go` | Entidades do domínio |
| `dto.go` | Objetos de entrada e saída da API |

Fluxo:

```
HTTP
 ↓
Handler
 ↓
Service
 ↓
Repository
 ↓
PostgreSQL
```

---

### `internal/user/`

Domínio responsável pelos usuários.

Segue a mesma organização:

- handler
- service
- repository
- model

---

### `internal/middleware/`

Middlewares compartilhados.

Exemplos:

- autenticação
- logger
- CORS
- rate limiter
- recovery

---

### `pkg/`

Código reutilizável.

Tudo aqui pode ser utilizado por vários domínios.

Exemplos:

```text
pkg/
├── base62/
├── hashid/
└── validator/
```

Esses pacotes normalmente não possuem dependência da aplicação.

---

### `scripts/`

Scripts auxiliares.

Exemplos:

- inicialização do ambiente;
- seed do banco;
- backup;
- deploy.

---

## Arquivos da raiz

### `.env`

Variáveis de ambiente.

---

### `go.mod`

Gerenciamento das dependências do projeto.

---

### `go.sum`

Checksums das dependências.

---

### `Makefile`

Automatiza comandos frequentes.

Exemplo:

```makefile
run:
	go run ./cmd/api

test:
	go test ./...

migrate:
	goose up

lint:
	golangci-lint run
```

---

# Fluxo da aplicação

```text
main.go
   │
   ▼
Config
   │
   ▼
Database
   │
   ▼
HTTP Server
   │
   ▼
Handler
   │
   ▼
Service
   │
   ▼
Repository
   │
   ▼
PostgreSQL
```

Essa estrutura separa claramente as responsabilidades da aplicação e facilita a manutenção e evolução do projeto.