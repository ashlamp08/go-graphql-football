package main

import (
	"encoding/json"
	"fmt"
	"github.com/ashlamp08/go-graphql-football/model"
	"github.com/graphql-go/graphql"
	"log"
)

var schema graphql.Schema

var aggregrateSchema = graphql.NewObject(graphql.ObjectConfig{Name: "RootQuery", Fields: graphql.Fields{
	"club": model.SingleClubSchema(),
	"list": model.ListClubSchema(),
}})

var aggregateMutations = graphql.NewObject(graphql.ObjectConfig{Name: "Mutation", Fields: graphql.Fields{
	"create_club":   model.CreateClubMutation(),
	"create_player": model.CreatePlayerMutation(),
}})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	return r
}

func init() {
	schemaConfig := graphql.SchemaConfig{
		Query:    aggregrateSchema,
		Mutation: aggregateMutations,
	}
	var err error
	schema, err = graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error : %v", err)
	}
}

func main() {

	querys := []string{` mutation {create_club(name:"Liverpool FC", location:"Liverpool") {name}}  `,
		`{ list {name, players{lastname}}}`,
		` mutation {create_player(firstname:"Prachet", lastname:"Sharma", position:"Forward", goals:7, playingclub:2) {name}}  `,
		`{ list {name, players{firstname, lastname, goals}}}`}

	for _, query := range querys {
		r := executeQuery(query, schema)
		rJSON, _ := json.Marshal(r)
		fmt.Printf("%s \n", rJSON)
	}
}
