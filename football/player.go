package football

type Player struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Position    string `json:"position"`
	Goals       int    `json:"goals"`
	PlayingClub int    `json:"playing_club"`
}
