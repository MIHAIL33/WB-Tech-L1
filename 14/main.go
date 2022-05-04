package main

/*Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.*/

import (
	"fmt"
	"reflect"
)

// Version 1 analog Version 2 (printf call reflect.TypeOf() )
func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// Version 2
func typeof_2(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// Version 3
func typeof_3(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan string:
		return "chan string"
	default: 
		return "unknown"
	}
}

func main() {
	//Version 1
	str := "data"
	fmt.Println(typeof(str)) //print string
	intt := 1
	fmt.Println(typeof(intt)) //print int

	//Version 2
	channel := make(chan string)
	fmt.Println(typeof_2(channel)) //print chan string

	b := true
	//Version 3
	fmt.Println("Version 3:")
	fmt.Println(typeof_3(b))
	fmt.Println(typeof_3(str))
	fmt.Println(typeof_3(intt))
	fmt.Println(typeof_3(channel))

}