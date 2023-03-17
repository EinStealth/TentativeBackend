package playername

import "github.com/EinStealth/TentativeBackend/internal/model"

// spacetimesにinsertする
func Post(name string) error {
	// dbmap初期化
	dbmap, err := model.InitDb()
	if err != nil {
		return err
	}
	defer dbmap.Db.Close()

	// dbmapにテーブル名登録
	dbmap.AddTableWithName(PlayerName{}, "playernames")

	// Insert
	player := &PlayerName{
		Name: name,
	}
	err = dbmap.Insert(player)
	if err != nil {
		return err
	}

	return nil
}
