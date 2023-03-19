package main

import "log"

/*
	1.	Дана структура Human (с произвольным набором полей и методов).
		Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

		Указывая структуре другую структуру как свойство, мы можем обращаться к методам вложенной структуры через родительскую.
*/

type Human struct {
	Name string
	Age  int
}

func (h *Human) GetName() string {
	return h.Name
}

func (h *Human) GetAge() int {
	return h.Age
}

func (h *Human) Say(s string) {
	log.Println(s)
}

type Action struct {
	Human
	Action string
	Age    int
}

func (a *Action) GetName() string {
	return "I have redefined my name!"
}

func main() {

	h := Human{
		Name: "Alan",
		Age:  23,
	}

	a := Action{
		Human:  h,
		Action: "dude",
		Age:    24,
	}
	log.Println(a.GetName())
	log.Println(a.GetAge())
	a.Say("I have inherited!")

}
