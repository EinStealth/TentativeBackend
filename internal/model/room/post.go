package room

import "github.com/EinStealth/TentativeBackend/internal/model"

// spacetimesにinsertする
func Post(secretWords string, isStart bool) error {
	// dbmap初期化
	dbmap, err := model.InitDb()
	if err != nil {
		return err
	}
	defer dbmap.Db.Close()

	// dbmapにテーブル名登録
	dbmap.AddTableWithName(Room{}, "room")

	// Insert
	room := &Room{
		SecretWords: secretWords,
		IsStart:     isStart,
	}
	err = dbmap.Insert(room)
	if err != nil {
		return err
	}

	return nil
}
