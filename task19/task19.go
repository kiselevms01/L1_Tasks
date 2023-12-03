package main

import "fmt"

// переворачиваем символы входной строки
func ReverseString(sourceString string) string {
	// Преобразовываем строку в слайс рун
	runeSlice := []rune(sourceString)

	for i := 0; i < len(runeSlice)/2; i++ {
		runeSlice[i], runeSlice[(len(runeSlice)-1)-i] = runeSlice[(len(runeSlice)-1)-i], runeSlice[i]
	}

	// Преобразовываем слайс рун в строку
	return string(runeSlice)
}

func main() {
	fmt.Println(ReverseString("самолет"))
	fmt.Println(ReverseString("ключ"))
	fmt.Println(ReverseString("сумка"))
}
