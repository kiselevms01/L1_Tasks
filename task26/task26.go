package main

import (
	"fmt"
	"strings"
)

// Функция проверки уникальности символов в строке
func CheckingUniqueness(sourceString string) bool {
	// Приводим к нижнему регистру
	sourceString = strings.ToLower(sourceString)

	// Преобразовываем строку в слайс рун
	runeSlice := []rune(sourceString)

	// Мапа для хранения встречающихся символов в строке
	mp := make(map[rune]struct{})

	for _, value := range runeSlice {
		_, ok := mp[value]
		if ok { // Ели символ value уже существует, то строка не уникальна
			return false
		} else {
			mp[value] = struct{}{}
		}
	}

	return true
}

func main() {
	fmt.Println(CheckingUniqueness("abcd"))
	fmt.Println(CheckingUniqueness("abCdefAaf"))
	fmt.Println(CheckingUniqueness("aabcd"))
	fmt.Println(CheckingUniqueness("abcdC"))
	fmt.Println(CheckingUniqueness("AbCdE"))
}
