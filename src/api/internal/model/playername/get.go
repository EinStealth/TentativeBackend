package playername

import "github.com/EinStealth/TentativeBackend/internal/model"

func GetByName(name string) ([]PlayerName, error) {
	// dbmap初期化
	dbmap, err := model.InitDb()
	if err != nil {
		return []PlayerName{}, err
	}
	defer dbmap.Db.Close()

	// secret_words指定でplayerのデータを取得
	var playerName []PlayerName
	_, err = dbmap.Select(&playerName, "SELECT * FROM playernames WHERE name=?", name)
	if err != nil {
		return []PlayerName{}, err
	}

	return playerName, nil
}