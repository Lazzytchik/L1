package main

import (
	"log"
	"math"
)

/*
	11. Реализовать пересечение двух неупорядоченных множеств.

	Для проверки уникальности значений мы используем map[T]struct{}
		Если значение есть в мапе то мы его запишем в итоговый массив.
		Мапа заполняется из первого слайса.
*/

func main() {
	a := []int{1, 2, 5, 7, 8, 9}
	b := []int{1, 3, 5, 10, 12}

	c := Intersect(a, b)

	log.Println(c)
}

func Intersect[T comparable](s1 []T, s2 []T) []T {
	//	Мы создаём слайс с capacity = минимальной длине из входных массивов, т.к. после итоговый слайс никогда не будет больше.
	//	Устанавливая правильное capacity мы можем избежать повторной аллокации памяти под массив при append.
	result := make([]T, 0, int(math.Min(float64(len(s1)), float64(len(s2)))))
	checker := make(map[T]struct{})

	for _, v := range s1 {
		checker[v] = struct{}{}
	}

	for _, v := range s2 {
		if _, ok := checker[v]; ok {
			checker[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}
