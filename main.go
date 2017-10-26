package main

import (
	"fmt"
	"joyfort/algorithm"
	"math/rand"
	"reflect"
	"time"
)

// import "algorithm"

const m_size = 9

//9X9 数独
var cells [9][9]int

//定义栈
var stack algorithm.Stack

//定义当前点
var coCurrent Coord

type Coord struct {
	x int
	y int
}

// 当前行走不通 ， 直接初始化 当前行
func initCells(a Coord) {
	for y := 0; y < m_size; y++ {
		cells[a.x][y] = 0
	}
}

//根据数独规则 ，产生当前点的可用数据
func getValidValueList(a Coord) [9]int {
	// fmt.Println(cells[a.x][a.y])
	// if cells[a.x][a.y] > 0 {
	//
	// }
	// var selectValue [9]int
	var leftValue [9]int
	// i := 0
	for i := 0; i < m_size; i++ {
		leftValue[i] = i + 1
	}

	for y := 0; y < m_size; y++ {
		if cells[a.x][y] > 0 {
			leftValue[cells[a.x][y]-1] = 0

		}
	}
	for x := 0; x < m_size; x++ {
		if cells[x][a.y] > 0 {
			leftValue[cells[x][a.y]-1] = 0

		}
	}
	// 3X3 矩阵
	// fmt.Println(leftValue)
	xstart := a.x - (a.x)%3
	ystart := a.y - (a.y)%3
	// fmt.Println(fmt.Sprintf(" 3X3 x %d y:%d", xstart, ystart))
	for ii := xstart; ii < xstart+3; ii++ {
		for oo := ystart; oo < ystart+3; oo++ {
			if cells[ii][oo] > 0 {
				leftValue[cells[ii][oo]-1] = 0
			}

		}
	}

	return leftValue

}

//下一个点
func NextCoord(a Coord) Coord {
	var b Coord
	b = Coord{a.x, a.y + 1}
	if a.x < m_size-1 && a.y == m_size-1 {
		b = Coord{a.x + 1, 0}
		stack.Push(b)
	}
	return b
	// }

}

//返回当前行的第一个点
func PrevCoord(a Coord) Coord {
	if a.y > 0 {
		initCells(a)
		return Coord{a.x, 0}
	}
	stack.Pop()
	// stack.Print()
	if stack.Empty() {
		return Coord{0, 0}
	}

	return stack.Top().(Coord)
}

func Generate_Randnum(param int) int {

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := rnd.Intn(param)

	return vcode
}

//根据可用数，随机取出来一个数当做当前点的数据
func randNum(num [9]int) int {
	var a [9]int
	var i int = 0

	for j := 0; j < m_size; j++ {
		if num[j] > 0 {
			a[i] = num[j]
			i++
		}
	}
	if len(a) == 0 {
		return 0
	}

	// r11 := RandInt64(0, i-1)
	r11 := Generate_Randnum(i)
	fmt.Println("select type", num, i, reflect.TypeOf(i), r11, a, "---------end")

	// fmt.Println(fmt.Printf("rand:range:%d,rand:%d;", len(a)-1, r11))

	return a[r11]
}

//数独主要函数
func generateValidMatrix() bool {
	// return false
	for {

		al := getValidValueList(coCurrent)

		valueLength := 0
		for i := 0; i < m_size; i++ {
			if al[i] > 0 {
				valueLength++
			}
		}
		if valueLength > 0 {

			cells[coCurrent.x][coCurrent.y] = randNum(al)

			if coCurrent.x == m_size-1 && coCurrent.y == m_size-1 {
				break
			}
			// else
			// {
			coCurrent = NextCoord(coCurrent)
			// }
		}
		if valueLength == 0 {

			if coCurrent.x == 0 && coCurrent.y == 0 {
				break
			}

			// cells[coCurrent.x][coCurrent.y] = 0
			coCurrent = PrevCoord(coCurrent)

		}
		// fmt.Println(cells)

	}
	printCells()
	return true
}

//output cells
func printCells() {
	for i := 0; i < m_size; i++ {
		fmt.Println(cells[i])
	}
	for i := 0; i < m_size; i++ {
		for j := 0; j < m_size; j++ {
			if Generate_Randnum(m_size) > 5 {
				// fmt.Print(cells[i][j], " ")
				cells[i][j] = 0
			} else {
				// fmt.Println("0", " ")
			}

		}
		// fmt.Println("")
	}
	fmt.Println("--------------")
	for i := 0; i < m_size; i++ {
		fmt.Println(cells[i])
	}
}

func main() {
	// fmt.Println(5 % 3)
	stack := algorithm.NewStack()
	coCurrent = Coord{0, 0}
	stack.Push(coCurrent)
	generateValidMatrix()
	// stack.Print()
	fmt.Println("gen success")
}
