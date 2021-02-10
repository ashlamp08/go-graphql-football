package model

import (
	"github.com/ashlamp08/gogql"
	"github.com/graphql-go/graphql"
)

type Player struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Position    string `json:"position"`
	Goals       int    `json:"goals"`
	PlayingClub int    `json:"playing_club"`
}

func SetupPlayerSchema(schemaBuilder *gogql.SchemaBuilder) *gogql.SchemaBuilder {

	schemaBuilder = schemaBuilder.AddMutationAction("create_player", "Create new player", Player{}, func(p graphql.ResolveParams) (interface{}, error) {
		player := Player{
			FirstName:   p.Args["first_name"].(string),
			LastName:    p.Args["last_name"].(string),
			Position:    p.Args["position"].(string),
			Goals:       p.Args["goals"].(int),
			PlayingClub: p.Args["playing_club"].(int),
		}
		for _, club := range Clubs {
			if player.PlayingClub == club.Id {
				club.Players = append(club.Players, player)
				Clubs[club.Id] = club
			}
		}
		return player, nil
	})

	return schemaBuilder
}
