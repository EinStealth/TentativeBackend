package room

import "github.com/EinStealth/TentativeBackend/internal/model"

func GetBySecretWords(secretWords string) ([]Room, error) {
	// dbmap初期化
	dbmap, err := model.InitDb()
	if err != nil {
		return []Room{}, err
	}
	defer dbmap.Db.Close()

	// secret_words指定でroomのデータを取得
	var room []Room
	_, err = dbmap.Select(&room, "SELECT * FROM room WHERE secret_words=?", secretWords)
	if err != nil {
		return []Room{}, err
	}

	return room, nil
}