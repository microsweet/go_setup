package main

import "fmt"

type Circle struct {
	radius float64
}

func main() {
	fmt.Println("Hello world!")

	var b, c int = 2, 3
	fmt.Println(b, c)

	var d int = 1
	add(&d)
	fmt.Println(d)

	gotoTag()

	/////////////swap/////////////////////
	var swapa int = 100
	var swapb int = 200

	fmt.Printf("交换前：a: %d;\tb:%d\n", swapa, swapb)

	swap(&swapa, &swapb)
	fmt.Printf("交换后：a: %d;\tb:%d\n", swapa, swapb)

	////////////////闭包///////////////////
	add_func := add_closure(1, 2)
	fmt.Println(add_func())
	fmt.Println(add_func())
	fmt.Println(add_func())

	//////////////圆面积//////////////
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积：", c1.getArea())

}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func add_closure(x1, x2 int) func() (int, int) {
	i := 1
	return func() (int, int) {
		i++
		return i, x1 + x2
	}

}

func swap(swapa, swapb *int) {
	var temp int
	temp = *swapa
	*swapa = *swapb
	*swapb = temp
}

func add(d *int) {
	*d = *d + 1

}

func gotoTag() {
	for m := 1; m < 10; m++ {
		n := 1
	Loop:
		if n <= m {
			fmt.Printf("%dx%d=%d\t", n, m, m*n)
			n++
			goto Loop
		} else {
			fmt.Println("")
		}
		n++
	}
}
