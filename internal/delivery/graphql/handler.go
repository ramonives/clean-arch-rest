package graphql

import (
	"net/http"

	"clean-arch-rest/internal/domain"

	"github.com/graphql-go/handler"
)

// NewGraphQLHandler cria um novo handler GraphQL
func NewGraphQLHandler(orderUseCase domain.OrderUseCase) http.Handler {
	schema := CreateSchema(orderUseCase)

	h := handler.New(&handler.Config{
		Schema:   schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return h
}
