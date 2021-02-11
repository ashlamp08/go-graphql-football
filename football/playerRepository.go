package football

import (
	"context"
)

func CreatePlayer(ctx context.Context, player Player) (result interface{}) {
	club := GetClubById(ctx, player.PlayingClub).(Club)
	club.Players = append(club.Players, player)
	return UpdateClub(ctx, club)
}
