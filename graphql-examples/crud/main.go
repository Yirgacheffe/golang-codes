package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

// product entity, use int64 instead of float64
type Product struct {
	ID    int64  `json:id`
	Name  string `json:name`
	Info  string `json:info,omitempty`
	Price int    `json:price`
}

var products = []Product{
	{
		ID:    1,
		Name:  "Chicha Morada",
		Info:  "Is a beverage originated by the actually national level",
		Price: 799,
	},
	{
		ID:    2,
		Name:  "Chicha de jora",
		Info:  "Corn beer chicha prepared by germinating maize.",
		Price: 595,
	},
	{
		ID:    3,
		Name:  "Pisco",
		Info:  "Is a colorless or yellowish-to-amber brandy produced",
		Price: 995,
	},
	{
		ID:    4,
		Name:  "Lego Ducati",
		Info:  "Serials of lego toy motor bike",
		Price: 1079,
	},
}

var Schema graphql.Schema

var productType = graphql.NewObject(graphql.ObjectConfig{
	Name: "product",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.Int},
		"name":  &graphql.Field{Type: graphql.String},
		"info":  &graphql.Field{Type: graphql.String},
		"price": &graphql.Field{Type: graphql.Int},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		/*
			Get (read) single product by id
			i.e. http://localhost:8080/graphql?query={product(id:1){name,info,price}}
		*/
		"product": &graphql.Field{
			Type:        productType,
			Description: "Get product by Id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if ok {
					for _, prod := range products {
						if int(prod.ID) == id {
							return prod, nil
						}
					}
				}
				return nil, nil // need refine the logic of return
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(productType),
			Description: "Get all product list",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return products, nil
			},
		},
	},
})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		/*
			product?query=mutation+_{create(name: "Inco Kola", info: "Inco Kola")}
		*/
		"create": &graphql.Field{
			Type:        productType,
			Description: "Create new product",
			Args: graphql.FieldConfigArgument{
				"name":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"info":  &graphql.ArgumentConfig{Type: graphql.String},
				"price": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				product := Product{
					ID:    int64(r.Intn(100000)),
					Name:  p.Args["name"].(string),
					Info:  p.Args["info"].(string),
					Price: p.Args["price"].(int),
				}
				products = append(products, product)
				return product, nil
			},
		},
		/*
			product?query=mutation+_{update(id:1,price:3.95){id,name,info,price}}
		*/
		"update": &graphql.Field{
			Type:        productType,
			Description: "Update product by id",
			Args: graphql.FieldConfigArgument{
				"id":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"name":  &graphql.ArgumentConfig{Type: graphql.String},
				"info":  &graphql.ArgumentConfig{Type: graphql.String},
				"price": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				name, nameOk := params.Args["name"].(string)
				info, infoOk := params.Args["info"].(string)
				price, priceOk := params.Args["price"].(int)

				product := Product{}

				for i, p := range products {
					if int64(id) == p.ID {
						if nameOk {
							products[i].Name = name
						}
						if infoOk {
							products[i].Info = info
						}
						if priceOk {
							products[i].Price = price
						}

						product = products[i]
						break
					}
				}

				return product, nil
			},
		},
		/* Delete product by id
		   product?query=mutation+_{delete(id:1){id,name,info,price}}
		*/
		"delete": &graphql.Field{
			Type:        productType,
			Description: "Delete product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				product := Product{}
				for i, p := range products {
					if int64(id) == p.ID {
						product = products[i]
						products = append(products[:i], products[i+1:]...)
					}
				}
				return product, nil
			},
		},
	},
})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(
		graphql.Params{
			Schema:        schema,
			RequestString: query,
		},
	)

	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func init() {
	s, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
		},
	)
	if err != nil {
		log.Fatalf("unable to register schema %v", err)
	}
	Schema = s
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), Schema)
		json.NewEncoder(w).Encode(result)
	})
	fmt.Println("Server is running on port: 8080")
	// 	fmt.Println("Test with Get              : curl -g 'http://localhost:8080/graphql?query={product{id,name}}'")
	http.ListenAndServe(":8080", nil)
}
