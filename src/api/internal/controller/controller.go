package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/EinStealth/TentativeBackend/internal/model/room"
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
		IsStart     bool   `json:"is_start"`
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
		IsStart     bool    `json:"is_start"`
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