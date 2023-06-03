package room

import (
	"github.com/EinStealth/TentativeBackend/internal/model"
)

// Roomのisstartを更新する
func UpdateIsStart(secret_words string, is_start int) error {
	// db初期化
	dbmap, err := model.InitDb()
	if err != nil {
		return err
	}
	defer dbmap.Db.Close()

	// dbからプレイヤーを取得
	roomData := Room{}
	err = dbmap.SelectOne(&roomData, "SELECT * FROM `rooms` WHERE `secret_words` = ?", secret_words)
	if err != nil {
		return err
	}

	// status更新
	roomData.IsStart = is_start
	dbmap.AddTableWithName(Room{}, "rooms").SetKeys(true, "SecretWords")
	_, err = dbmap.Update(&roomData)
	if err != nil {
		return err
	}

	return nil
}

// Roomのisstartを更新する
func UpdateDeamon(secret_words string, deamon int) error {
	// db初期化
	dbmap, err := model.InitDb()
	if err != nil {
		return err
	}
	defer dbmap.Db.Close()

	// dbからプレイヤーを取得
	roomData := Room{}
	err = dbmap.SelectOne(&roomData, "SELECT * FROM `rooms` WHERE `secret_words` = ?", secret_words)
	if err != nil {
		return err
	}

	// status更新
	roomData.Deamon = deamon
	dbmap.AddTableWithName(Room{}, "rooms").SetKeys(true, "SecretWords")
	_, err = dbmap.Update(&roomData)
	if err != nil {
		return err
	}

	return nil
}
