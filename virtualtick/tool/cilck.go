package tool

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

//打印鼠标当前位置（记得用协程，不然会卡）
func PrintPositionFor()  {
	for {
		x, y := robotgo.GetMousePos()
		fmt.Println(`x：`, x, ` y：`, y)
		time.Sleep(time.Second)
	}
}

//丝滑的移动
func Move2(x,y int)  {
	robotgo.MoveMouseSmooth(x, y,0.1,0.1)
	fmt.Printf("鼠标移动到了（%d,%d）", x, y)
}

func PrintPos()  {
	x, y := robotgo.GetMousePos()
	fmt.Println(`x：`, x, ` y：`, y)
	time.Sleep(time.Second)
}

