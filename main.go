package main

import (
	"./gql"
	"./models"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"net/http"
)

var zradla = []*models.Zradlo{
	{1, "Pizza", 600},
	{2, "Sushi", 1000},
	{3, "Borsch", 200},
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(
		graphql.Params{
			Schema:        schema,
			RequestString: query,
		},
	)

	return result
}

func main() {
	var schema, _ = gql.Schema(&zradla)
	http.HandleFunc("/zradlo", func(writer http.ResponseWriter, request *http.Request) {
		result := executeQuery(request.URL.Query().Get("query"), schema)
		_ = json.NewEncoder(writer).Encode(result)

	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
