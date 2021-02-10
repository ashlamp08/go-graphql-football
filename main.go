package main

import (
	"encoding/json"
	"fmt"
	"github.com/ashlamp08/go-graphql-football/model"
	"github.com/ashlamp08/gogql"
	"github.com/graphql-go/graphql"
	"log"
)

var schema graphql.Schema

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	return r
}

func init() {
	schemaBuilder := gogql.NewSchemaBuilder()
	schemaBuilder = model.SetupClubSchema(schemaBuilder)
	schemaBuilder = model.SetupPlayerSchema(schemaBuilder)
	schema = schemaBuilder.Build()
}

func main() {

	querys := []string{` mutation {create_club(name:"Liverpool FC", location:"Liverpool") {name}}  `,
		`{ list {name, players{last_name}}}`,
		` mutation {create_player(first_name:"Prachet", last_name:"Sharma", position:"Forward", goals:7, playing_club:2) {first_name}}  `,
		`{ list {location, players{first_name, goals}}}`}

	for _, query := range querys {
		r := executeQuery(query, schema)
		rJSON, _ := json.Marshal(r)
		fmt.Printf("%s \n", rJSON)
	}
}
