package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type CustomID struct {
	value string
}

func (id *CustomID) String() string {
	return id.value
}

func NewCustomID(v string) *CustomID {
	return &CustomID{value: v}
}

var CustomScalarType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "CustomScalarType",
	Description: "The `CustomScalarType` scalar type represents an ID Object",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case CustomID:
			return value.String()
		case *CustomID:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			return NewCustomID(value)
		case *string:
			return NewCustomID(*value)
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return NewCustomID(valueAST.Value)
		default:
			return nil
		}
	},
})

type Customer struct {
	ID *CustomID `json:"id"`
}

var CustomerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Customer",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: CustomerScalarType,
		},
	},
})

func main() {
	schema, err := graghql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"customers": &graphql.Field{
					Type: graphql.NewList(CustomerType),
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: CustomScalarType},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// id := p.Args["id"]
						log.Printf("id from arguments: %+v", id)
						customers := []Customer{
							Customer{ID: NewCustomID("fb8378783f")},
						}
						return customers, nil
					},
				},
			},
		}),
	})

	if err != nil {
		log.Fatal(err)
	}

	query := `
		query{
			customers {
				id
			}
		}
	`

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
		VariableValues: map[string]interface{}{
			"id": "7367748fsb3",
		},
	})

	if len(result.Errors) > 0 {
		log.Fatal(result)
	}

	b, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
