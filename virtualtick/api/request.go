package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ReqeustBasicGameInfo()  {
	resp, err := http.Get("http://127.0.0.1:21337/positional-rectangles")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	b2 := &BasicGameInfo{}
	err2 := json.Unmarshal(body, b2)
	if err != nil {
		fmt.Println(err2)
	}
	//当改变游戏模式时
	if BasicInfoAPI.GameState != b2.GameState{
		fmt.Printf("游戏模式改变:%s-->%s\n",BasicInfoAPI.GameState,b2.GameState)
	}
	//当有牌且在游戏中：
	if len(NowRectanglesAPI)!=0&&b2.GameState=="InProgress"{
		//b2.PrintCards()
	}
	BasicInfoAPI = *b2
	NowRectanglesAPI = b2.Rectangles
}

func (b *BasicGameInfo)RequestNewInfo() *BasicGameInfo {
	resp, err := http.Get("http://127.0.0.1:21337/positional-rectangles")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	b2 := &BasicGameInfo{}
	err2 := json.Unmarshal(body, b2)
	if err != nil {
		fmt.Println(err2)
	}
	BasicInfoAPI = *b2
	NowRectanglesAPI = b2.Rectangles
	return b2
}

//场上所有明面上的牌
func (b *BasicGameInfo)PrintCards()  {
	for i, i2 := range b.Rectangles {
		fmt.Printf("第%d张,%v\n",i+1,i2)
	}
}

//我的所有牌
func (b *BasicGameInfo)PrintMyCards()  {
	i:=0
	for _, i2 := range b.Rectangles {
		if i2.LocalPlayer == true&& (i2.CardCode != "face"){
			fmt.Printf("第%d张,%v\n",i+1,i2)
			i++
		}
	}
}

//获取场上对手的牌
func (b *BasicGameInfo) PrintEnemyDeck()  {
	i:=0
	for _, i2 := range b.Rectangles {
		if (i2.LocalPlayer == false) && (i2.CardCode != "face"){
			fmt.Printf("第%d张,%v\n",i+1,i2)
			i++
		}
	}
}

//场上随从
func (b *BasicGameInfo)GetMyUnit(){
	i:=0
	for _, i2 := range b.Rectangles {
		if i2.LocalPlayer == true&& (i2.CardCode != "face"){
			if i2.TopLeftY > 100 {
				fmt.Printf("第%d张,%v\n",i+1,i2)
				i++
			}
		}
	}
}

//场上随从数量
func (b *BasicGameInfo)GetMyUnitNumber() int  {
	var num int
	for _, i2 := range b.Rectangles {
		if i2.LocalPlayer == true&& (i2.CardCode != "face"){
			if i2.TopLeftY > 100 {
				num ++
			}
		}
	}
	return num
}

//手牌
func (b *BasicGameInfo)GetMyHandCard() {
	i:=0
	for _, i2 := range b.Rectangles {
		if i2.LocalPlayer == true&& (i2.CardCode != "face"){
			if i2.TopLeftY <= 100 {
				fmt.Printf("第%d张,%v\n",i+1,i2)
				i++
			}
		}
	}
}

//手牌数量
func (b *BasicGameInfo)GetMyHandCardNumber() int  {
	var num int
	for _, i2 := range b.Rectangles {
		if i2.LocalPlayer == true&& (i2.CardCode != "face"){
			if i2.TopLeftY <= 100 {
				num ++
			}
		}
	}
	return num
}

//手牌里第i张牌的位置
func (b *BasicGameInfo)GetOneCardPosition (num int) (int,int) {
	i:=0
	for _, i2 := range b.Rectangles {
		if i2.LocalPlayer == true&& (i2.CardCode != "face"){
			if i2.TopLeftY <= 100 {
				fmt.Printf("第%d张,%v\n",i+1,i2)
				i++
			}
			if i-1 == num{
				return i2.TopLeftX,1080-i2.TopLeftY
			}
		}
	}
	return 523,1030
}

