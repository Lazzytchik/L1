package main

import "fmt"

/*
	14. Что выведет данная программа и почему?

	Вывод [b b a]\n[a a].
	Причина аналогична ответу из 13 вопроса.
	При append локальный slice стал ссылаться на новый массив и все изменения были применены к нему.
	Изначальный slice остался неизменен.
*/

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Println(slice)
	}(slice)
	fmt.Println(slice)
}
