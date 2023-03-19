package main

import (
	"log"
	"math/big"
)

/*
	22. Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a, b, значение которых > 2^20.

	Мы используем пакет big для этого.
*/

func main() {
	var a big.Float
	var b big.Float
	var c big.Float

	a.SetString("3982636272982939773687768723")
	b.SetString("28200272635728982652587878")
	log.Printf("Result of +: %s", c.Add(&a, &b).String())
	log.Printf("Result of -: %s", c.Sub(&a, &b).String())
	log.Printf("Result of /: %s", c.Quo(&a, &b).String())
	log.Printf("Result of *: %s", c.Mul(&a, &b).String())
}
