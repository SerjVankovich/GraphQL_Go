package main

import (
	"./db"
	"./gql"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"net/http"
)

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
	dataBase := db.Connect()
	defer dataBase.Close()
	var schema, _ = gql.Schema(dataBase)
	http.HandleFunc("/zradlo", func(writer http.ResponseWriter, request *http.Request) {
		result := executeQuery(request.URL.Query().Get("query"), schema)
		_ = json.NewEncoder(writer).Encode(result)

	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
