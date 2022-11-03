package oop

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

// structs Point类型定义一个方法
func (p Point) Distance(q Point) float64 {
	// 求斜边. 勾股定律: 勾3股4弦5
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Add(another Point) Point {
	return Point{p.X + another.X, p.Y + another.Y}
}

func (p Point) Sub(another Point) Point {
	return Point{p.X - another.X, p.Y - another.Y}
}

func (p Point) Print() {
	fmt.Printf("{%f, %f}\n", p.X, p.Y)
}

func DistanceExample() {

	p := Point{1, 2}
	q := Point{4, 6}

	distanceFormP := p.Distance   // 方法值(相当于C语言的函数地址,函数指针)
	fmt.Println(distanceFormP(q)) // "5"
	fmt.Println(p.Distance(q))    // "5"

	distanceFormQ := q.Distance   //实际上distanceFormP 就绑定了 p接收器的方法Distance
	fmt.Println(distanceFormQ(p)) // "5"
	fmt.Println(q.Distance(p))    // "5"
}

// time.AfterFunc这个函数的功能是在指定的延迟时间之后来执行一个(译注：另外的)函数。且这个函数操作的是一个Rocket对象r
