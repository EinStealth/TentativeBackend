package room

import "github.com/EinStealth/TentativeBackend/internal/model"

// spacetimesにinsertする
func Post(secret_words string, is_start bool) error {
	// dbmap初期化
	dbmap, err := model.InitDb()
	if err != nil {
		return err
	}
	defer dbmap.Db.Close()

	// dbmapにテーブル名登録
	dbmap.AddTableWithName(Room{}, "rooms")

	// Insert
	rooms := &Room{
		SecretWords: secret_words,
		IsStart:     is_start,
	}
	err = dbmap.Insert(rooms)
	if err != nil {
		return err
	}

	return nil
}
