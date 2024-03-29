package main

import (
	"log"
	"strconv"
)

/*
	21. Реализовать паттерн «адаптер» на любом примере.

	Две структуры Int и String e обоих есть метод GetValue()
	И структура StringReverser которая принимает только тип Stringer (не путать со стандартным) определяющий метод GetValue() string.
	Однако у Int есть лишь метод GetValue() int.
	Поэтому мы сделаем адаптер IntAdapter который будет реализовывать метод GetValue() string, конвертируя int в string с помощью strconv.Itoa()
	Таким образом мы можем отдать в качестве параметра адаптер вместо Int в StringReverser.Reverse()
*/

func main() {

	str := String{"Hello"}
	num := Int{482936}

	numAdapter := IntAdapter{&num}

	reverser := StringReverser{}

	log.Println(reverser.Reverse(str))
	log.Println(reverser.Reverse(numAdapter))

}

type Stringer interface {
	GetValue() string
}

type StringReverser struct {
}

func (s *StringReverser) Reverse(str Stringer) string {
	runes := []rune(str.GetValue())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type Int struct {
	Value int
}

func (i Int) GetValue() int {
	return i.Value
}

type String struct {
	Value string
}

func (s String) GetValue() string {
	return s.Value
}

type IntAdapter struct {
	Int *Int
}

func (ia IntAdapter) GetValue() string {
	return strconv.Itoa(ia.Int.GetValue())
}
