package lor

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"math/rand"
	"time"
	"virtualtick/api"
)

//都是在1920x1080下

type Position struct {
	x int
	y int
}


var (
	Setting =  Position{
		x: 1833,
		y: 49,
	}//设置
	Pass =  Position{
		x: 1689,
		y: 527,
	}//结束回合
	Giveup = Position{
		x: 827,
		y: 904,
	}//投降
	Conti = Position{
		x: 1268,
		y: 989,
	}//结束之后的继续
	Play = Position{
		x: 80,
		y: 370,
	}//左边的游玩
	Fightwiithplayer = Position{
		x: 300,
		y: 255,
	}//玩家对战
	Cards = Position{
		x: 639,
		y: 309,
	}//卡组
	Simple = Position{
		x: 1760,
		y: 174,
	}//一般对战
	Point = Position{
		x: 1268,
		y: 989,
	}//积分对战
	StartGame = Position{
		x: 1639,
		y: 974,
	}//开始对战
	OK = Position{
		x: 1046,
		y: 596,
	}//开始对战
	Attack1 = Position{
		x: 1487,
		y:893,
	}
	Attack2 = Position{
		x: 1487,
		y:674,
	}

)

const (
	color_setting = "f0b178"
	color_contin = "c27f3b"
)


//加工一个位置为附近随机点
func (p *Position)RadonPosition() *Position {
	return &Position{
		x: p.x + rand.Intn(5),
		y: p.y + rand.Intn(5),
	}
}

//移动并点击
func (p *Position)MoveAndClick()  *Position{
	robotgo.MoveMouseSmooth(p.x, p.y,0.1,0.1)
	robotgo.MouseClick(`left`, false)
	fmt.Printf("鼠标移动到了（%d,%d）并点击了一次\n", p.x, p.y)
	return p
}

//封装
func (p *Position)Touch()  *Position {
	p.RadonPosition().MoveAndClick()
	//暂停1s
	time.Sleep(time.Second)
	return p
}

func (p *Position)Pass() *Position {
	robotgo.KeyTap("space")
	time.Sleep(time.Second)
	return p
}

func (p *Position)JustMove()  *Position{
	robotgo.MoveMouseSmooth(p.x, p.y,0.1,0.1)
	return p
}

func (p *Position)Down()  *Position{
	robotgo.MouseToggle("down")
	return p
}

func (p *Position)Up()  *Position{
	robotgo.MouseToggle("up")
	return p
}

func (p *Position)Drag(position Position)  *Position{
	robotgo.MoveSmooth(position.x, position.y,1.1,0.8)
	return p
}

func (p *Position)Sleep() *Position {
	time.Sleep(time.Millisecond*500)
	return p
}

func StartInit()  {
	//先check一下是否进入准备阶段

	//点击游玩
	Play.Touch()

	//点击与玩家对战
	Fightwiithplayer.Touch()

	//点击第一套卡牌
	Cards.Touch()

	//一般对战
	Simple.Touch()

	//开始游戏
	StartGame.Touch()

	MainGame(900)

}

func MainGameWithoutStart()  {
	//开始游戏
	StartGame.Touch()

	MainGame(900)
}

var flag int

func MainGame(n int)  {
	for {
		Pass.Pass()
		fmt.Println(flag)
		if flag<=n{
			flag ++
			if Continue()==1{
				time.Sleep(time.Second*2)
				//开始游戏
				StartGame.Touch()
				flag = 0
				continue
			}
		}else {
			if Continue()==1{
				time.Sleep(time.Second*2)
				//开始游戏
				StartGame.Touch()
				flag = 0
				continue
			}
			if GiveUp() == 1{
				//投降完了的界面
				//点击继续
				time.Sleep(time.Second*7)
				Conti.Touch()
				time.Sleep(time.Second*2)
				//开始游戏
				StartGame.Touch()
				flag = 0
			}
		}
		time.Sleep(time.Second)
	}
}


//获得对应位置的颜色
func (p *Position)GetColorOf()string  {
	return robotgo.GetPixelColor(p.x, p.y)
}

//投降
func GiveUp() int {
	setting := Setting.GetColorOf()
	if  setting == color_setting{
		Setting.Touch()
		Giveup.Touch()
		OK.Touch()
		return 1
	}else {
		time.Sleep(time.Second)
		fmt.Println("bad color",setting)
		return 0
	}
}


func Continue() int {
	cont := Conti.GetColorOf()
	if cont == color_contin{
		Conti.Touch()
		return 1
	}else {
		time.Sleep(time.Second)
		return 0
	}
}


//交互部分
func UserUI()  {


	fmt.Println("---从主界面按s进入游戏---")
	robotgo.EventHook(hook.KeyDown, []string{"s"}, func(e hook.Event) {
		fmt.Println("s")
		StartInit()
	})

	fmt.Println("---按d开始游戏---")
	robotgo.EventHook(hook.KeyDown, []string{"d"}, func(e hook.Event) {
		fmt.Println("d")
		MainGameWithoutStart()
	})

	fmt.Println("---v---")
	robotgo.EventHook(hook.KeyDown, []string{"v"}, func(e hook.Event) {
		fmt.Println("v")
		GetPosition()
	})

	fmt.Println("---2---")
	robotgo.EventHook(hook.KeyDown, []string{"2"}, func(e hook.Event) {
		fmt.Println("2")
		var b api.BasicGameInfo
		unit := b.RequestNewInfo().GetMyUnitNumber()
		fmt.Println(unit)
	})

	fmt.Println("---h---")
	robotgo.EventHook(hook.KeyDown, []string{"h"}, func(e hook.Event) {
		fmt.Println("h")
		var b api.BasicGameInfo
		b.RequestNewInfo().GetMyHandCard()

	})

	fmt.Println("---a---")
	robotgo.EventHook(hook.KeyDown, []string{"a"}, func(e hook.Event) {
		fmt.Println("a")
		var a Attack
		a.AutoAttack()
	})

	fmt.Println("---m---")
	robotgo.EventHook(hook.KeyDown, []string{"m"}, func(e hook.Event) {
		fmt.Println("m")
		var pg PlayG
		pg.PlayARandCard()
	})


	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func GetColor()  {
	x, y := robotgo.GetMousePos()
	pos := &Position{
		x: x,
		y: y,
	}
	color := robotgo.GetPixelColor(pos.x, pos.y)
	fmt.Println("color----", color)
}

func GetPosition()  {
	x, y := robotgo.GetMousePos()
	pos := &Position{
		x: x,
		y: y,
	}
	fmt.Println("pos----",pos)
}
