package main

import (
	"github.com/ashlamp08/go-graphql-football/football"
	"github.com/ashlamp08/go-graphql-football/infrastructure"
	"github.com/ashlamp08/gogql"
	"github.com/go-chi/chi"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
	"net/url"
)

var schema graphql.Schema

func init() {
	// initialize schema for GQL
	schemaBuilder := gogql.NewSchemaBuilder()
	schemaBuilder = football.SetupClubSchema(schemaBuilder)
	schemaBuilder = football.SetupPlayerSchema(schemaBuilder)
	schema = schemaBuilder.Build()

	// initialize environment
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Nepal")
	env := infrastructure.Environment{}
	env.SetEnvironment()
	env.LoadConfig()
	env.InitMongoDB()
}

func main() {
	routes := chi.NewRouter()
	r := RegisterRoutes(routes)
	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
