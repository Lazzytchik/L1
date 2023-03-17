package main

import (
	"log"
	"math"
)

/*
	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
	Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножествах не важна.

	Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

*/

func main() {

	temperatures := []float32{-25.0, -27.0, 13.0, -21.0, 15.5, 24.5, 19.0, 32.5, -5.0}
	clastered := make(map[int][]float32)

	for _, temperature := range temperatures {
		label := int(math.Ceil(float64(temperature))) / 10 * 10
		v, ok := clastered[label]
		if !ok {
			clastered[label] = make([]float32, 0, 10)
		}

		clastered[label] = append(v, temperature)

	}

	log.Println(clastered)

}