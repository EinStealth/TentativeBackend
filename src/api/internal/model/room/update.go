package room

import (
	"github.com/EinStealth/TentativeBackend/internal/model"
)

// Roomのisstartを更新する
func UpdateIsStart(id, is_start int) error {
	// db初期化
	dbmap, err := model.InitDb()
	if err != nil {
		return err
	}
	defer dbmap.Db.Close()

	// dbからプレイヤーを取得
	roomData := Room{}
	err = dbmap.SelectOne(&roomData, "SELECT * FROM `rooms` WHERE `id` = ?", id)
	if err != nil {
		return err
	}

	// status更新
	roomData.IsStart = is_start
	dbmap.AddTableWithName(Room{}, "rooms").SetKeys(true, "Id")
	_, err = dbmap.Update(&roomData)
	if err != nil {
		return err
	}

	return nil
}
