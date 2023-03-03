package player

type Player struct {
	Id           int     `db:"id"`
	SecretWords  string  `db:"secret_words"`
	Name         string  `db:"name"`
	Icon         int     `db:"icon"`
	Status       int     `db:"status"`
}