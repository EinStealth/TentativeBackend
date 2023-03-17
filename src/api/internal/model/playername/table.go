package playername

type PlayerName struct {
	Id           int     `db:"id"`
	Name         string  `db:"name"`
}