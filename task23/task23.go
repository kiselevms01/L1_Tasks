package main

import "fmt"

func Delete(data []string, i int) []string {
	// удаляем элемент слайса с помощью append
	return append(data[:i], data[i+1:]...)
}

func main() {
	fullSlice := []string{"Blue", "Yellow", "Green", "Grey", "Black", "White"}

	var i int

	fmt.Println("Введите id удаляемого элемента:")

	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("What?")
	}

	fmt.Printf("Result: %v\n", Delete(fullSlice, i))
}
