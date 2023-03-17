package main

import (
	"log"
	"math"
)

/*
	11. Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	a := []int{1, 2, 5, 7, 8, 9}
	b := []int{1, 3, 5, 10, 12}

	c := Intersect(a, b)

	log.Println(c)
}

func Intersect[T comparable](s1 []T, s2 []T) []T {
	result := make([]T, 0, int(math.Max(float64(len(s1)), float64(len(s2)))))
	checker := make(map[T]struct{})

	for _, v := range s1 {
		if _, ok := checker[v]; !ok {
			checker[v] = struct{}{}
			//result = append(result, v)
		}
	}

	for _, v := range s2 {
		if _, ok := checker[v]; ok {
			checker[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}
