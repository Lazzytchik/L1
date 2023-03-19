package main

import (
	"log"
	"strings"
)

/*
	15. К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

	Если пользоваться много раз функцией somefunc(), то у нас забьётся память.
	Так как строка это слайс, а значит под капотом создаётся массив.
	Строка большая, однако мы возвращаем лишь слайс её первых 100 элементов.
	Однако сам массив при этом остаётся в памяти, так как на него всё ещё есть ссылка.
	В итоге мы пользуемся лишь малой частью того под что выделили память.

	Выход такой: Вместо возврата слайса от созданной строки, мы вернём копию этого слайса.
	Таким образом, память под строку освободится и у нас останется нужная её часть.
*/

var justString string

func main() {
	log.Println(justString)
	someFuncUpdated()
	log.Println(justString, len(justString))
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func someFuncUpdated() {
	v := createHugeString(1 << 10)
	//justString = strings.Clone(v[:100])
	b := make([]byte, 100)
	copy(b, v[:100])
	justString = string(b)
}

func createHugeString(size int) string {
	var sb strings.Builder

	for i := 0; i < size; i++ {
		sb.WriteString("s")
	}

	return sb.String()
}
