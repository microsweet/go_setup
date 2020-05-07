// hello world
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/shopspring/decimal"
)

// Circle 圆
type Circle struct {
	radius float64
}

func main() {
	testdeci()
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
	addFunc := addClosure(1, 2)
	fmt.Println(addFunc())
	fmt.Println(addFunc())
	fmt.Println(addFunc())

	//////////////圆面积//////////////
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积：", c1.getArea())

	/////////////数组//////////////
	var balance = [5]int{1, 2, 3, 4, 5}
	getAverage(balance[:], 5)

	/////////////类型初始值///////////
	var (
		a1 int
		a2 string
		a3 float32
		a4 bool
	)
	fmt.Printf("int: ")
	fmt.Println(a1)
	fmt.Printf("string: ")
	fmt.Println(a2)
	fmt.Printf("float32: ")
	fmt.Println(a3)
	fmt.Printf("boolean: ")
	fmt.Println(a4)
	///////////变量交换///////
	var jh1 int = 1
	var jh2 int = 2
	jh1 = jh1 ^ jh2
	jh2 = jh2 ^ jh1
	jh1 = jh1 ^ jh2
	fmt.Println(jh1)
	fmt.Println(jh2)

	//////////////sin///////////////
	sin()

}

func getAverage(arr []int, size int) {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)
	fmt.Println(avg)

}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func addClosure(x1, x2 int) func() (int, int) {
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

func sin() {
	// 图片大小
	const size = 300
	// 根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	// 遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// 填充为白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	// 从0到最大像素生成x坐标
	for x := 0; x < size; x++ {
		// 让sin的值的范围在0~2Pi之间
		s := float64(x) * 2 * math.Pi / size
		// sin的幅度为一半的像素。向下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2
		// 用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}
	// 创建文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	// 使用png格式将数据写入文件
	png.Encode(file, pic) //将image信息写入文件中
	// 关闭文件
	file.Close()
}
func testdeci() {

	x := 74.96
	y := 20.48
	fmt.Println(x - y)
	xd := decimal.NewFromFloat(x)
	yd := decimal.NewFromFloat(y)
	c := xd.Sub(yd)
	cc := c.String()
	fmt.Println(xd.Sub(yd), cc)
}
