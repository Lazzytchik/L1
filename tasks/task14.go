package main

import (
	"fmt"
	"log"
	"reflect"
)

/*
	14. Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

	Тут 3 варианта как мы можем получить тип из переменной.
*/

func main() {
	log.Println(GetTypeString(2))
	log.Println(GetTypeReflect(2).String())
	log.Println(GetTypeManually(2))
}

func GetTypeString(obj interface{}) string {
	return fmt.Sprintf("%T", obj)
}

func GetTypeReflect(obj interface{}) reflect.Type {
	return reflect.TypeOf(obj)
}

func GetTypeManually(obj interface{}) string {
	switch obj.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	default:
		return "unknown"
	}
}
