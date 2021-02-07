package model

import "github.com/graphql-go/graphql"

type Club struct {
	Id       int
	Name     string
	Location string
	Players  []Player
}

var clubs map[int]Club

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

	clubs = map[int]Club{}
	clubs[1] = *club
}

func CreateClubMutation() *graphql.Field {
	return &graphql.Field{
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
				Id:       len(clubs) + 1,
				Name:     p.Args["name"].(string),
				Location: p.Args["location"].(string),
				Players:  []Player{},
			}
			clubs[club.Id] = club
			return club, nil
		},
	}
}

func SingleClubSchema() *graphql.Field {
	return &graphql.Field{
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
	}
}

func ListClubSchema() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(clubType),
		Description: "Get Club List",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var l []Club
			for _, v := range clubs {
				l = append(l, v)
			}
			return l, nil
		},
	}
}
