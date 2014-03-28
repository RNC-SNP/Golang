package main

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	GREEN
)

type Color byte //type可用于定义任何类型，相当于C语言的typedef

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box

func (c Color) String() string { //可为任何类型定义method
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "GREEN"}
	return strings[c]
}

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) { //指针method
	b.color = c //这里不用*b，会自动转换
}

func (bl BoxList) BiggestsColor() Color {
	v := 0.00
	c := Color(BLACK)
	for _, b := range bl {
		if b.Volume() > v {
			v = b.Volume()
			c = b.color
		}
	}
	return c
}

func (bl BoxList) PaintAll(c Color) {
	for i, _ := range bl {
		bl[i].SetColor(c) //这里bl不用取地址，会自动转换为指针
	}
}

func main() {
	boxes := BoxList{
		Box{1, 2, 3, RED},
		Box{2, 3, 4, BLACK},
		Box{3, 4, 5, RED},
	}
	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume())
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestsColor().String())
	boxes.PaintAll(BLACK)
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestsColor().String())
}
