package main

import (
	"log"
	"math"
)

/*
	24.	Разработать программу нахождения расстояния между двумя точками,
		которые представлены в виде структуры Point с инкапсулированными параметрами x, y и конструктором.

	Чтобы инкапсулировать параметры мы напишем их в нижнем регистре.
	Чтобы инициализировать Point, мы сделаем конструктор NewPoint таким образом мы можем создать точку с координатами, но не можем изменить координаты после инициализации.
*/

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Magnitude(t *Point) float64 {
	return math.Sqrt(math.Pow(p.x-t.x, 2) + math.Pow(p.y-t.y, 2))
}

func main() {
	first := NewPoint(6, 8)
	second := NewPoint(3, 4)

	log.Println(first.Magnitude(second))
}
