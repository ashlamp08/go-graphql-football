package football

type Club struct {
	Id       int      `bson:"_id" json:"id"`
	Name     string   `bson:"name" json:"name"`
	Location string   `bson:"location" json:"location"`
	Players  []Player `bson:"players" json:"players"`
}
