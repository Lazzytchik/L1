package main

import (
	"log"
	"time"
)

/*
	2. Что такое интерфейсы, как они применяются?

	Интерфейс - это сущность описывающая поведение.
	В Go выражается в типе interface который содержит сигнатуры методов.
	Если тип реализует методы интерфейса, то переменная типа интерфейса может хранить данный тип.

	Интерфейсы позволяют использовать разные реализации для одного функционала, если можно определить одинаковый интерфейс использования.
	(пример с люстрой, лампочками и цоколем)

	Интерфейсы также удобны когда ты хочешь написать расширяемый функционал, чтобы даже при смене реализации остальной функционал продолжал работать с новой реализацией.
*/

type Human interface {
	Sleep()
	Talk(s string)
}

type Dude struct {
	Name string
	Age  int
}

func (d *Dude) Talk(s string) {
	log.Println("Hey there!", s)
}

func (d *Dude) Sleep() {
	log.Println("Start sleeping...")
	time.Sleep(5 * time.Second)
	log.Println("I'm awake!")
}

func main() {
	var obj Human

	obj = &Dude{"Karl", 23}

	obj.Talk("Come here!")
	obj.Sleep()
}
