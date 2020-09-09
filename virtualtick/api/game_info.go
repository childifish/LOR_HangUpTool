package api

import (
	"time"
)

type BasicGameInfo struct {
	PlayerName string `json:"PlayerName"`
	OpponentName string `json:"OpponentName"`
	GameState string `json:"GameState"`			//有InProgress和Menus
	Screen Screen `json:"Screen"`				//屏幕尺寸结构体
	Rectangles []Rectangles `json:"Rectangles"` //卡牌信息结构体
}
type Screen struct {
	ScreenWidth int `json:"ScreenWidth"`
	ScreenHeight int `json:"ScreenHeight"`
}
type Rectangles struct {
	CardID int `json:"CardID"`
	CardCode string `json:"CardCode"`
	TopLeftX int `json:"TopLeftX"`
	TopLeftY int `json:"TopLeftY"`
	Width int `json:"Width"`
	Height int `json:"Height"`
	LocalPlayer bool `json:"LocalPlayer"`
}

/*
{"PlayerName":"xxx","OpponentName":"decks_mediumzed_name","GameState":"InProgress","Screen":{"ScreenWidth":1920,"ScreenHeight":1080},"Rectangles":[{"CardID":1791786086,"CardCode":"face","TopLeftX":179,"TopLeftY":481,"Width":117,"Height":117,"LocalPlayer":true},{"CardID":1625469617,"CardCode":"face","TopLeftX":179,"TopLeftY":716,"Width":117,"Height":117,"LocalPlayer":false},{"CardID":1810562329,"CardCode":"02NX001","TopLeftX":478,"TopLeftY":77,"Width":175,"Height":245,"LocalPlayer":true},{"CardID":191548564,"CardCode":"01NX004","TopLeftX":634,"TopLeftY":84,"Width":176,"Height":246,"LocalPlayer":true},{"CardID":405904526,"CardCode":"02BW038","TopLeftX":792,"TopLeftY":86,"Width":176,"Height":247,"LocalPlayer":true},{"CardID":2116406318,"CardCode":"02NX008","TopLeftX":950,"TopLeftY":84,"Width":177,"Height":248,"LocalPlayer":true},{"CardID":743424477,"CardCode":"01IO005","TopLeftX":817,"TopLeftY":980,"Width":127,"Height":160,"LocalPlayer":false},{"CardID":1664183267,"CardCode":"01IO045","TopLeftX":974,"TopLeftY":980,"Width":127,"Height":160,"LocalPlayer":false},{"CardID":1669047327,"CardCode":"02NX003","TopLeftX":1110,"TopLeftY":78,"Width":178,"Height":249,"LocalPlayer":true},{"CardID":1806990324,"CardCode":"01NX046","TopLeftX":1271,"TopLeftY":67,"Width":179,"Height":250,"LocalPlayer":true}]}
*/

var (
	BasicInfoAPI BasicGameInfo
	NowRectanglesAPI []Rectangles
)


func GetInfoFromGame()  {
	for {
		ReqeustBasicGameInfo()
		time.Sleep(time.Second)
	}
}

