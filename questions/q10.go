package main

import "fmt"

/*
	10. Что выведет данная программа и почему?

	В Go всё по умолчанию передаётся по значению. В том числе и указатели.
	При передаче в функцию update параметра указателя создался новый указатель на тот же участок памяти.
	Соответственно изменение этого указателя внутри функции бессмысленно так как этот указатель будет только внутри функции.
	Переданный указатель не изменится, следовательно, вывод будет 1 \n 1

	Однако если изменить значение на которое ссылается указатель, то, т.к. оба указателя ссылаются на одно место в памяти изменение хранимого значения отразится на обоих указателях.
*/

func update(p *int) {
	b := 2
	p = &b
}

func change(p *int) {
	*p = 2
}

func main() {
	var (
		a = 1
		p = &a
	)

	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
	change(p)
	fmt.Println(*p)
}
