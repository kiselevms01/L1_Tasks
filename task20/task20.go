package main

import (
	"fmt"
	"strings"
)

func Reverse(str string) string {
	sliceRev := strings.Split(str, " ") // создаем слайс стрингов

	left, right := 0, len(sliceRev)-1
	for ; left < right; left, right = left+1, right-1 {
		sliceRev[left], sliceRev[right] = sliceRev[right], sliceRev[left]
	}

	return strings.Join(sliceRev, " ") // объединям перевернутые слова в одну строку
}

func main() {
	fmt.Println(Reverse("Собака - лучший друг человека"))
	fmt.Println(Reverse("Какой сегодня день недели?"))
}
