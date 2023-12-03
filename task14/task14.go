package main

import (
	"fmt"
	"reflect"
)

// с помощью метода TypeOf пакета "reflect"
func TypeVar(v interface{}) {
	fmt.Printf("%v is %v", v, reflect.TypeOf(v))
}

// C использованием флага %T в пакете fmt, чтобы получить представление типа
/*
func TypeVar(v interface{}) {
	xType := fmt.Sprintf("%v : %T", v, v)
	fmt.Println(xType)
}
*/

func main() {
	var x interface{} = "23"
	TypeVar(x)
}
