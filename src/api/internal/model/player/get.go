package player

import "github.com/EinStealth/TentativeBackend/internal/model"

func GetBySecretWords(secretWords string) ([]Player, error) {
	// dbmap初期化
	dbmap, err := model.InitDb()
	if err != nil {
		return []Player{}, err
	}
	defer dbmap.Db.Close()

	// secret_words指定でplayerのデータを取得
	var player []Player
	_, err = dbmap.Select(&player, "SELECT * FROM players WHERE secret_words=?", secretWords)
	if err != nil {
		return []Player{}, err
	}

	return player, nil
}