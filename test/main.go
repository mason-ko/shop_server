package main

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql-go-handler"
	"net/http"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"x": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: nil,
					Description:  "",
				},
				"y": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: nil,
					Description:  "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello World!", nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	// ...
	Query: queryType,
})

func main() {
	// Schema
	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	// and serve!
	http.ListenAndServe(":8080", nil)
}
