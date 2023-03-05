package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/EinStealth/TentativeBackend/internal/model/room"
	"github.com/EinStealth/TentativeBackend/internal/model/player"
	"github.com/EinStealth/TentativeBackend/internal/model/location"
)

// @Summary     test api
// @Description ping pong
// @Accept      json
// @Success     200 {string} string "pong"
// @Router      /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// @Summary 指定された合言葉の(ステータス)を取得する
// @Param   time query string true "time"
// @Accept  json
// @Produce json
// @Success 200 {array} controller.GetRoom.JsonResponse
// @Filure  400 {string} string err.Error()
// @Filure  500 {string} string err.Error()
// @Router  /api/v1/room [get]
func GetRoom(c *gin.Context) {
	// param取得
	secretWords := c.Query("secret_words")

	// 指定した合言葉のステータスを取得
	res, err := room.GetBySecretWords(secretWords)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 受け取ったものをJsonResponseに詰め直す
	type JsonResponse struct {
		SecretWords string `json:"secret_words"`
		IsStart     int    `json:"is_start"`
	}
	jsonRes := make([]JsonResponse, len(res))
	for i := 0; i < len(res); i++ {
		jsonRes[i].SecretWords = res[i].SecretWords
		jsonRes[i].IsStart = res[i].IsStart
	}
	c.JSON(200, jsonRes)
}

// @Summary 指定された合言葉の(ステータス)を格納するAPI
// @Param   request body controller.PostRoom.JsonRequest true "request json"
// @Accept  json
// @Produce application/json
// @Success 200 {string} string "success"
// @Filure  400 {string} string err.Error()
// @Filure  500 {string} string err.Error()
// @Router  /api/v1/room [post]
func PostRoom(c *gin.Context) {
	// param取得
	type JsonRequest struct {
		SecretWords string  `json:"secret_words"`
		IsStart     int     `json:"is_start"`
	}
	var req JsonRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// roomに受け取ったパラメータをpostする
	err = room.Post(req.SecretWords, req.IsStart)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// @Summary プレイヤーの状態を更新するAPI
// @Param   id     path int true "player id"
// @Param   status path int true "player status"
// @Accept  json
// @Produce json
// @Success 200 {string} string "success"
// @Filure  400 {string} string "引数が違います"
// @Filure  500 {string} string err.Error()
// @Router  /api/v1/players/{id}/status/{status} [post]
func UpdateRoomIsStart(c *gin.Context) {
	// パスパラメータ取得
	secret_words := c.Param("secret_words")
	is_start, err := strconv.Atoi(c.Param("is_start"))
	if err != nil || is_start < 0 || 1 < is_start {
		c.JSON(400, gin.H{
			"message": "is_startが違います",
		})
		return
	}

	// isstart更新
	err = room.UpdateIsStart(secret_words, is_start)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

// @Summary 指定された合言葉の(ステータス)を取得する
// @Param   time query string true "time"
// @Accept  json
// @Produce json
// @Success 200 {array} controller.GetRoom.JsonResponse
// @Filure  400 {string} string err.Error()
// @Filure  500 {string} string err.Error()
// @Router  /api/v1/room [get]
func GetPlayer(c *gin.Context) {
	// param取得
	secretWords := c.Query("secret_words")

	// 指定した合言葉のステータスを取得
	res, err := player.GetBySecretWords(secretWords)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 受け取ったものをJsonResponseに詰め直す
	type JsonResponse struct {
		SecretWords  string  `json:"secret_words"`
		Name         string  `json:"name"`
		Icon         int     `json:"icon"`
		Status       int     `json:"status"`
	}
	jsonRes := make([]JsonResponse, len(res))
	for i := 0; i < len(res); i++ {
		jsonRes[i].SecretWords = res[i].SecretWords
		jsonRes[i].Name = res[i].Name
		jsonRes[i].Icon = res[i].Icon
		jsonRes[i].Status = res[i].Status
	}
	c.JSON(200, jsonRes)
}

// @Summary 指定された合言葉の(ステータス)を格納するAPI
// @Param   request body controller.PostRoom.JsonRequest true "request json"
// @Accept  json
// @Produce application/json
// @Success 200 {string} string "success"
// @Filure  400 {string} string err.Error()
// @Filure  500 {string} string err.Error()
// @Router  /api/v1/room [post]
func PostPlayer(c *gin.Context) {
	// param取得
	type JsonRequest struct {
		SecretWords  string  `json:"secret_words"`
		Name         string  `json:"name"`
		Icon         int     `json:"icon"`
		Status       int     `json:"status"`
	}
	var req JsonRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// playerに受け取ったパラメータをpostする
	err = player.Post(req.SecretWords, req.Name, req.Icon, req.Status)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// @Summary プレイヤーの状態を更新するAPI
// @Param   id     path int true "player id"
// @Param   status path int true "player status"
// @Accept  json
// @Produce json
// @Success 200 {string} string "success"
// @Filure  400 {string} string "引数が違います"
// @Filure  500 {string} string err.Error()
// @Router  /api/v1/players/{id}/status/{status} [post]
func UpdatePlayerStatus(c *gin.Context) {
	// パスパラメータ取得
	name := c.Param("name")
	status, err := strconv.Atoi(c.Param("status"))
	if err != nil || status < 1 || 4 < status {
		c.JSON(400, gin.H{
			"message": "statusが違います",
		})
		return
	}

	// status更新
	err = player.UpdateStatus(name, status)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

// @Summary 指定された合言葉の(ステータス)を取得する
// @Param   time query string true "time"
// @Accept  json
// @Produce json
// @Success 200 {array} controller.GetRoom.JsonResponse
// @Filure  400 {string} string err.Error()
// @Filure  500 {string} string err.Error()
// @Router  /api/v1/room [get]
func GetLocation(c *gin.Context) {
	// param取得
	secretWords := c.Query("secret_words")
	relativeTime := c.Query("relative_time")

	// 指定した合言葉のステータスを取得
	res, err := location.GetBySecretWordsAndSpacetime(secretWords, relativeTime)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 受け取ったものをJsonResponseに詰め直す
	type JsonResponse struct {
		SecretWords  string  `json:"secret_words"`
		RelativeTime string  `json:"relative_time"`
		Latitude     float64 `json:"latitude"`
		Longitude    float64 `json:"longitude"`
		Status       int     `json:"status"`
	}
	jsonRes := make([]JsonResponse, len(res))
	for i := 0; i < len(res); i++ {
		jsonRes[i].SecretWords = res[i].SecretWords
		jsonRes[i].RelativeTime = res[i].RelativeTime
		jsonRes[i].Latitude = res[i].Latitude
		jsonRes[i].Longitude = res[i].Longitude
		jsonRes[i].Status = res[i].Status
	}
	c.JSON(200, jsonRes)
}

// @Summary 指定された合言葉の(ステータス)を格納するAPI
// @Param   request body controller.PostRoom.JsonRequest true "request json"
// @Accept  json
// @Produce application/json
// @Success 200 {string} string "success"
// @Filure  400 {string} string err.Error()
// @Filure  500 {string} string err.Error()
// @Router  /api/v1/room [post]
func PostLocation(c *gin.Context) {
	// param取得
	type JsonRequest struct {
		SecretWords  string  `json:"secret_words"`
		RelativeTime string  `json:"relative_time"`
		Latitude     float64 `json:"latitude"`
		Longitude    float64 `json:"longitude"`
		Status       int     `json:"status"`
	}
	var req JsonRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// playerに受け取ったパラメータをpostする
	err = location.Post(req.SecretWords, req.RelativeTime, req.Latitude, req.Longitude, req.Status)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}
