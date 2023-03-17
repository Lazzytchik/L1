package main

import "log"

/*
	12. Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {

	s := []string{"cat", "dude", "cat", "cat", "dude", "man", "cat"}

	log.Println(MakeSet(s))

}

func MakeSet(s []string) []string {
	result := make([]string, 0, len(s))
	known := make(map[string]struct{})

	for _, v := range s {
		if _, ok := known[v]; !ok {
			result = append(result, v)
			known[v] = struct{}{}
		}
	}

	return result
}
