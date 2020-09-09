package main

import (
	"virtualtick/api"
	"virtualtick/lor"
)

func main()  {
	go api.GetInfoFromGame()
	lor.UserUI()
}



