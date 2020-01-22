package leaderboard

import "github.com/adrenallen/leaderboarder/datasource"

//Handler takes care of retrieving and saving
//leaderboard entries using a datasource
type Handler struct {
	Data datasource.DataSource
}

func (h Handler) Test() {
	return
}

func (h Handler) NewEntry(name string, score float64, meta string) error{
	e := datasource.Entry{Name: name, Score: score, Meta: meta}
	h.Data.SaveNew(e)
	return nil
}