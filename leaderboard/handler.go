package leaderboard

import "github.com/adrenallen/leaderboarder/datasource"

//Handler takes care of retrieving and saving
//leaderboard entries using a datasource
type Handler struct {
	ds *datasource.DataSource
}

func (h Handler) Test() {
	return
}