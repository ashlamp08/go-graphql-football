package football

import (
	"context"
	"github.com/ashlamp08/gogql"
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
)

func SetupPlayerSchema(schemaBuilder *gogql.SchemaBuilder) *gogql.SchemaBuilder {
	schemaBuilder = schemaBuilder.AddMutationAction("create_player", "Create new player", Player{}, createPlayerResolver)
	return schemaBuilder
}

var createPlayerResolver = func(p graphql.ResolveParams) (interface{}, error) {
	var player Player
	mapstructure.Decode(p.Args, &player)
	CreatePlayer(context.Background(), player)
	return player, nil
}
