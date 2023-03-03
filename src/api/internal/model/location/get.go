package location

import "github.com/EinStealth/TentativeBackend/internal/model"

func GetBySecretWordsAndSpacetime(secret_words string, relative_time string) ([]Location, error) {
	// dbmap初期化
	dbmap, err := model.InitDb()
	if err != nil {
		return []Location{}, err
	}
	defer dbmap.Db.Close()

	// secret_words指定でplayerのデータを取得
	var location []Location
	_, err = dbmap.Select(&location, "SELECT * FROM locations WHERE secret_words=? AND relative_time=?", secret_words, relative_time)
	if err != nil {
		return []Location{}, err
	}

	return location, nil
}