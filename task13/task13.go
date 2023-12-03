package main

import "fmt"

// Сложение и вычитание
func Swap(a, b *int) {
	*a += *b
	*b = *a - *b
	*a = *a - *b
}

// Двойное присваивание
/*
func Swap(a, b *int) {
	*a, *b = *b, *a
}
*/

func main() {
	a, b := 2, 8

	fmt.Printf("old : a = %d, b = %d\n", a, b) // начальные значения

	Swap(&a, &b)

	fmt.Printf("new: : a = %d, b = %d\n", a, b) // после изменения
}
