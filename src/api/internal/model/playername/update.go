package playername

import (
	"github.com/EinStealth/TentativeBackend/internal/model"
)

// Playerのstatusを更新する
func UpdatePlayerName(preName string, newName string) error {
	// db初期化
	dbmap, err := model.InitDb()
	if err != nil {
		return err
	}
	defer dbmap.Db.Close()

	// dbからプレイヤーを取得
	playerData := PlayerName{}
	err = dbmap.SelectOne(&playerData, "SELECT * FROM `playernames` WHERE `name` = ?", preName)
	if err != nil {
		return err
	}

	// status更新
	playerData.Name = newName
	dbmap.AddTableWithName(PlayerName{}, "playernames").SetKeys(true, "Name")
	_, err = dbmap.Update(&playerData)
	if err != nil {
		return err
	}

	return nil
}
