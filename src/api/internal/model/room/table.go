package room

type Room struct {
	Id           int     `db:"id"`
	SecretWords  string  `db:"secret_words"`
	IsStart      int    `db:"is_start"`
}