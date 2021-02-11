module github.com/ashlamp08/go-graphql-football

go 1.15

require (
	github.com/ashlamp08/gogql v0.0.0-20210210160618-c0cadd586c01
	github.com/go-chi/chi v1.5.2
	github.com/graphql-go/graphql v0.7.9
	github.com/graphql-go/handler v0.2.3
	github.com/mitchellh/mapstructure v1.4.1
	go.mongodb.org/mongo-driver v1.4.6
	gopkg.in/yaml.v2 v2.4.0
)

//replace github.com/ashlamp08/gogql => ../gogql
