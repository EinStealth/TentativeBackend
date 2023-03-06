package location

type Location struct {
	Id           int     `db:"id"`
	SecretWords  string  `db:"secret_words"`
	RelativeTime string  `db:"relative_time"`
	Latitude     float64 `db:"latitude"`
	Longitude    float64 `db:"longitude"`
	// 鬼か逃げか: 鬼: 1, 逃げ: 2
	// 人か罠か: 人: 1, 罠: 2、罠削除: 3
	// 2桁で送る
	Status       int     `db:"status"`
}