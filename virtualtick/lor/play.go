package lor

import (
	"math/rand"
	"virtualtick/api"
)

type Attack struct {
}
type PlayG struct {
}

func (a *Attack)Select()  {
}

//随从位置
var unitSlice  = []int{961, 876, 793, 724, 650, 599}

//自动攻击
func (a *Attack)AutoAttack()  {
	var b api.BasicGameInfo
	unit := b.RequestNewInfo().GetMyUnitNumber()
	a.Attack(unit)
}

//攻击
func (a *Attack)Attack(n int)  {
	if n <1 {
		return
	}
	if n >5{
		n = 5
	}
	p := Position{
		x: unitSlice[n-1],
		y: 900,
	}
	p.JustMove().Sleep()
	p.Down().Sleep().Drag(Attack1).Sleep().Drag(Attack2).Sleep().Up().Sleep()
	Pass.Pass()
}

//出一张随机牌
func  (pg *PlayG)PlayARandCard()  {
	pg.GetHandCardPos().PlayCard()
}

//获取一张牌的位置
func (pg *PlayG)GetHandCardPos() Position {
	var b api.BasicGameInfo
	number := b.RequestNewInfo().GetMyHandCardNumber()
	x,y := b.RequestNewInfo().GetOneCardPosition(rand.Intn(number))
	//因为给的位置是左上角,直接用点不到
	p := Position{
		x:x+50,
		y:y+10,
	}
	return p
}

//出牌
func (p Position)PlayCard()  {
	p.JustMove().Sleep()
	p.Down().Sleep().Drag(Position{x: 900+rand.Intn(70), y: 500+rand.Intn(50)}).Sleep().Up().Sleep()
}
