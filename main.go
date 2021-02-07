package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
)

type Club struct {
	Id       int
	Name     string
	Location string
	Players  []Player
}

type Player struct {
	FirstName   string
	LastName    string
	Position    string
	Goals       int
	PlayingClub int
}

func populate() []Club {
	player := &Player{
		FirstName:   "Mason",
		LastName:    "Mount",
		Position:    "Midfield",
		Goals:       5,
		PlayingClub: 1,
	}

	club := &Club{
		Id:       1,
		Name:     "Chelsea FC",
		Location: "London",
		Players:  []Player{*player},
	}

	return []Club{*club}
}

func main() {
	// GraphQL Objects
	var playerType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
			"firstname": &graphql.Field{
				Type: graphql.String,
			},
			"lastname": &graphql.Field{
				Type: graphql.String,
			},
			"position": &graphql.Field{
				Type: graphql.String,
			},
			"goals": &graphql.Field{
				Type: graphql.Int,
			},
			"playingclub": &graphql.Field{
				Type: graphql.Int,
			},
		}})

	var clubType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Club",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"location": &graphql.Field{
				Type: graphql.String,
			},
			"players": &graphql.Field{
				Type: graphql.NewList(playerType),
			},
		},
	})

	clubs := populate()

	var mutationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Type:        clubType,
				Description: "Create a new Club",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"location": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					club := Club{
						Name:     p.Args["name"].(string),
						Location: p.Args["location"].(string),
						Players:  []Player{},
					}
					clubs = append(clubs, club)
					return club, nil
				},
			},
		},
	})

	// Schema
	fields := graphql.Fields{
		"club": &graphql.Field{
			Type:        clubType,
			Description: "Get Club by Id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// take in the ID argument
				id, ok := p.Args["id"].(int)
				if ok {
					// Parse our club array or matching id
					for _, club := range clubs {
						if int(club.Id) == id {
							return club, nil
						}
					}
				}
				return nil, nil
			},
		},

		"list": &graphql.Field{
			Type:        graphql.NewList(clubType),
			Description: "Get Club List",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return clubs, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: mutationType}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error : %v", err)
	}

	querys := []string{` { mutation {create(name:"Liverpool FC", location:"Liverpool") {name}}  }`,
		`{ list {name, players{lastname}}}`}

	for _, query := range querys {
		params := graphql.Params{Schema: schema, RequestString: query}
		r := graphql.Do(params)
		if len(r.Errors) > 0 {
			log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
		}
		rJSON, _ := json.Marshal(r)
		fmt.Printf("%s \n", rJSON)
	}
}
