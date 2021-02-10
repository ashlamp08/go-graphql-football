package model

import (
	"github.com/ashlamp08/gogql"
	"github.com/graphql-go/graphql"
)

type Club struct {
	Id       int      `json:"id" unique:"true"`
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Players  []Player `json:"players"`
}

var Clubs map[int]Club

func init() {
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

	Clubs = map[int]Club{}
	Clubs[1] = *club
}

func SetupClubSchema(schemaBuilder *gogql.SchemaBuilder) *gogql.SchemaBuilder {

	schemaBuilder = schemaBuilder.AddQueryAction("club", "Get club by Id", Club{}, func(p graphql.ResolveParams) (interface{}, error) {
		// take in the ID argument
		id, ok := p.Args["id"].(int)
		if ok {
			// Parse our club array or matching id
			for _, club := range Clubs {
				if int(club.Id) == id {
					return club, nil
				}
			}
		}
		return nil, nil
	})

	schemaBuilder = schemaBuilder.AddQueryAction("list", "Get club list", []Club{}, func(p graphql.ResolveParams) (interface{}, error) {
		var l []Club
		for _, v := range Clubs {
			l = append(l, v)
		}
		return l, nil
	})

	schemaBuilder = schemaBuilder.AddMutationAction("create_club", "Create a club", Club{}, func(p graphql.ResolveParams) (interface{}, error) {
		club := Club{
			Id:       len(Clubs) + 1,
			Name:     p.Args["name"].(string),
			Location: p.Args["location"].(string),
			Players:  []Player{},
		}
		Clubs[club.Id] = club
		return club, nil
	})

	return schemaBuilder
}