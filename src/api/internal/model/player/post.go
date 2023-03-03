package player

import "github.com/EinStealth/TentativeBackend/internal/model"

// spacetimesにinsertする
func Post(secret_words string, name string, icon int, status int) error {
	// dbmap初期化
	dbmap, err := model.InitDb()
	if err != nil {
		return err
	}
	defer dbmap.Db.Close()

	// dbmapにテーブル名登録
	dbmap.AddTableWithName(Player{}, "players")

	// Insert
	player := &Player{
		SecretWords: secret_words,
		Name: name,
		Icon: icon,
		Status: status,
	}
	err = dbmap.Insert(player)
	if err != nil {
		return err
	}

	return nil
}
