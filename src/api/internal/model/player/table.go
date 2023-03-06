package player

type Player struct {
	Id           int     `db:"id"`
	SecretWords  string  `db:"secret_words"`
	Name         string  `db:"name"`
	Icon         int     `db:"icon"`
	// 鬼か逃げか: 鬼: 1, 逃げ: 2
	// 普通: 1, トラップにかかった: 2, 捕まった: 3, 逃げ切った: 4
	// 2桁で送る
	Status       int     `db:"status"`
}