package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

var Schema graphql.Schema

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"me": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Context.Value("current_user"), nil
				},
			},
		},
	},
)

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	user := struct {
		ID   int    `json:"id"`
		Name string `name:"name"`
	}{1, "Feel Free User"}

	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: r.URL.Query().Get("query"),
		Context:       context.WithValue(context.Background(), "current_user", user),
	})

	if len(result.Errors) > 0 {
		log.Printf("wrong result, got some errors: %v", result.Errors)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func init() {
	s, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

	if err != nil {
		log.Fatalf("failed to registered schema, error: %v", err)
	}
	Schema = s
}

func main() {
	http.HandleFunc("/graphql", graphqlHandler)
	fmt.Println("Running on port 8080")
	fmt.Println("Test with Get		: curl -g 'http://localhost:8080/graphql?query={me{id,name}}'")
	http.ListenAndServe(":8080", nil)
}
