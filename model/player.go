package model

import "github.com/graphql-go/graphql"

type Player struct {
	FirstName   string
	LastName    string
	Position    string
	Goals       int
	PlayingClub int
}

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
	},
})

func CreatePlayerMutation() *graphql.Field {
	return &graphql.Field{
		Type:        clubType,
		Description: "Create a new Player",
		Args: graphql.FieldConfigArgument{
			"firstname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lastname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"position": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"goals": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"playingclub": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			player := Player{
				FirstName:   p.Args["firstname"].(string),
				LastName:    p.Args["lastname"].(string),
				Position:    p.Args["position"].(string),
				Goals:       p.Args["goals"].(int),
				PlayingClub: p.Args["playingclub"].(int),
			}
			for _, club := range clubs {
				if player.PlayingClub == club.Id {
					club.Players = append(club.Players, player)
					clubs[club.Id] = club
				}
			}
			return player, nil
		},
	}
}
