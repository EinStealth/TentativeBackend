package location

import "github.com/EinStealth/TentativeBackend/internal/model"

// spacetimesにinsertする
func Post(secret_words string, relative_time string, latitude float64, longitude float64, status int) error {
	// dbmap初期化
	dbmap, err := model.InitDb()
	if err != nil {
		return err
	}
	defer dbmap.Db.Close()

	// dbmapにテーブル名登録
	dbmap.AddTableWithName(Location{}, "locations")

	// Insert
	location := &Location{
		SecretWords: secret_words,
		RelativeTime: relative_time,
		Latitude: latitude,
		Longitude: longitude,
		Status: status,
	}
	err = dbmap.Insert(location)
	if err != nil {
		return err
	}

	return nil
}
