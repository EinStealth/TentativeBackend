package player

import (
	"github.com/EinStealth/TentativeBackend/internal/model"
)

// Playerのstatusを更新する
func UpdateStatus(name string, status int) error {
	// db初期化
	dbmap, err := model.InitDb()
	if err != nil {
		return err
	}
	defer dbmap.Db.Close()

	// dbからプレイヤーを取得
	playerData := Player{}
	err = dbmap.SelectOne(&playerData, "SELECT * FROM `players` WHERE `name` = ?", name)
	if err != nil {
		return err
	}

	// status更新
	playerData.Status = status
	dbmap.AddTableWithName(Player{}, "players").SetKeys(true, "Name")
	_, err = dbmap.Update(&playerData)
	if err != nil {
		return err
	}

	return nil
}
