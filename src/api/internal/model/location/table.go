package location

type Location struct {
	Id           int     `db:"id"`
	SecretWords  string  `db:"secret_words"`
	RelativeTime string  `db:"relative_time"`
	Latitude     float64 `db:"latitude"`
	Longitude    float64 `db:"longitude"`
	Status       int     `db:"status"`
}