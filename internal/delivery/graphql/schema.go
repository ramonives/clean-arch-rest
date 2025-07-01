package graphql

import (
	"clean-arch-rest/internal/domain"

	"github.com/graphql-go/graphql"
)

// OrderType define o tipo GraphQL para Order
var OrderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Order",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"customer_id": &graphql.Field{
			Type: graphql.Int,
		},
		"total": &graphql.Field{
			Type: graphql.Float,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// CreateSchema cria o schema GraphQL com as queries e mutations
func CreateSchema(orderUseCase domain.OrderUseCase) *graphql.Schema {
	fields := graphql.Fields{
		"listOrders": &graphql.Field{
			Type: graphql.NewList(OrderType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return orderUseCase.ListOrders()
			},
		},
		"createOrder": &graphql.Field{
			Type: OrderType,
			Args: graphql.FieldConfigArgument{
				"customer_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"total": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
				"status": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				customerID := uint(p.Args["customer_id"].(int))
				total := p.Args["total"].(float64)
				status := p.Args["status"].(string)

				order := &domain.Order{
					CustomerID: customerID,
					Total:      total,
					Status:     status,
				}

				err := orderUseCase.CreateOrder(order)
				if err != nil {
					return nil, err
				}

				return order, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	rootMutation := graphql.ObjectConfig{Name: "RootMutation", Fields: fields}
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		panic(err)
	}

	return &schema
}
