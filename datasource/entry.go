package datasource

// Entry is a leaderboard row containing
// a name, score, and metadata
type Entry struct {
	Name string
	Score float64

	//meta data about this entry
	Meta interface{}
}