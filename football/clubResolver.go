package football

import (
	"context"
	"github.com/ashlamp08/gogql"
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
)

func SetupClubSchema(schemaBuilder *gogql.SchemaBuilder) *gogql.SchemaBuilder {
	schemaBuilder = schemaBuilder.AddQueryAction("club", "Get club by Id", Club{}, getClubByIdResolver)
	schemaBuilder = schemaBuilder.AddQueryAction("list", "Get club list", []Club{}, getClubListResolver)
	schemaBuilder = schemaBuilder.AddMutationAction("create_club", "Create a club", Club{}, createClubResolver)
	return schemaBuilder
}

var createClubResolver = func(p graphql.ResolveParams) (interface{}, error) {
	var club Club
	mapstructure.Decode(p.Args, &club)
	club.Players = []Player{}
	CreateClub(context.TODO(), club)
	return club, nil
}

var getClubByIdResolver = func(p graphql.ResolveParams) (interface{}, error) {
	// take in the ID argument
	id, ok := p.Args["id"].(int)
	if ok {
		return GetClubById(context.Background(), id), nil
	}
	return nil, nil
}

var getClubListResolver = func(p graphql.ResolveParams) (interface{}, error) {
	return GetClubList(context.Background(), 10), nil
}
