# Clean Architecture REST API

Este projeto implementa uma API REST com Clean Architecture em Go, incluindo endpoints REST, gRPC e GraphQL para gerenciamento de orders.

## 🎯 Desafio Implementado

Este projeto foi desenvolvido para completar o desafio de criar um usecase de listagem de orders com:
- ✅ **Endpoint REST** (GET /order)
- ✅ **Service ListOrders com gRPC**
- ✅ **Query ListOrders GraphQL**
- ✅ **Migrações automáticas** do banco de dados
- ✅ **Docker Compose** para inicialização completa
- ✅ **Arquivo api.http** com requests para testar

## 🚀 Inicialização Automática

**IMPORTANTE:** Ao executar `docker compose up`, tanto as migrações quanto a aplicação são inicializadas automaticamente!

### Como Executar

1. **Certifique-se de que o Docker Desktop está rodando**

2. **Execute o comando:**
   ```bash
   docker compose up
   ```

3. **Pronto!** O sistema irá:
   - 🗄️ Subir o PostgreSQL na porta 5432
   - 📊 Executar as migrações automaticamente (criação da tabela + dados de exemplo)
   - 🔄 Aguardar o banco estar pronto
   - 🚀 Inicializar a aplicação Go com todos os serviços
   - 🌐 Disponibilizar todas as APIs nas portas configuradas

### Execução em Background

```bash
docker compose up -d
```

### Parar os Serviços

```bash
docker compose down
```

## 🛠️ Tecnologias Utilizadas

- Go 1.21
- PostgreSQL
- Docker & Docker Compose
- Gin (HTTP Framework)
- gRPC
- GraphQL
- GORM (ORM)

## 📁 Estrutura do Projeto

```
clean-arch-rest/
├── internal/
│   ├── domain/           # Entidades e interfaces
│   ├── usecase/          # Casos de uso
│   ├── infrastructure/   # Implementações externas
│   └── delivery/         # Controllers e handlers
├── proto/               # Arquivos protobuf
├── migrations/          # Migrações do banco
├── docker-compose.yaml  # Configuração Docker
├── Dockerfile          # Build da aplicação
└── main.go             # Arquivo principal
```

## 🌐 Portas dos Serviços

- **REST API:** http://localhost:8080
- **GraphQL:** http://localhost:8081/graphql
- **gRPC:** localhost:9090
- **PostgreSQL:** localhost:5432

## 📋 Endpoints Disponíveis

### REST API

- `POST /order/` - Criar uma nova order
- `GET /order/` - Listar todas as orders

### GraphQL

- Query: `listOrders` - Listar todas as orders
- Mutation: `createOrder` - Criar uma nova order

### gRPC

- `CreateOrder` - Criar uma nova order
- `ListOrders` - Listar todas as orders

## 🧪 Testando a API

### Usando o arquivo api.http

O projeto inclui o arquivo `api.http` com requests prontos para testar:

```http
### Criar uma nova order
POST http://localhost:8080/order/
Content-Type: application/json

{
    "customer_id": 1,
    "total": 150.50,
    "status": "pending"
}

### Listar todas as orders
GET http://localhost:8080/order/
```

### Dados de Exemplo

O banco de dados já vem populado com 3 orders de exemplo:

```json
[
  {
    "id": 1,
    "customer_id": 1,
    "total": 150.50,
    "status": "pending",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  },
  {
    "id": 2,
    "customer_id": 2,
    "total": 299.99,
    "status": "completed",
    "created_at": "2024-01-01T11:00:00Z",
    "updated_at": "2024-01-01T11:00:00Z"
  },
  {
    "id": 3,
    "customer_id": 3,
    "total": 99.90,
    "status": "shipped",
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
]
```

### Testando REST

**Listar orders existentes:**
```bash
curl http://localhost:8080/order/
```

**Criar uma nova order:**
```bash
curl -X POST http://localhost:8080/order/ \
  -H "Content-Type: application/json" \
  -d '{"customer_id": 4, "total": 199.99, "status": "pending"}'
```

### Testando GraphQL

Acesse http://localhost:8081/graphql e use a query:

```graphql
query {
  listOrders {
    id
    customer_id
    total
    status
    created_at
    updated_at
  }
}
```

### Testando gRPC

Use uma ferramenta como [grpcurl](https://github.com/fullstorydev/grpcurl) para testar o serviço gRPC na porta 9090.

## ⚙️ Configuração do Banco

- **Host:** localhost (ou postgres no Docker)
- **Porta:** 5432
- **Usuário:** ramon
- **Senha:** 1234
- **Database:** clean_arch_db

## 📊 Estrutura da Tabela Orders

```sql
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER NOT NULL,
    total DECIMAL(10,2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 🔧 Desenvolvimento Local

Se preferir executar localmente para desenvolvimento:

1. **Subir apenas o banco de dados:**
   ```bash
   docker compose up postgres -d
   ```

2. **Instalar dependências:**
   ```bash
   go mod tidy
   ```

3. **Executar a aplicação:**
   ```bash
   go run main.go
   ```