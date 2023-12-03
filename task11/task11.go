package main

import "fmt"

// Intersection - ищет пересечения двух неупорядоченных множеств
func Intersection(list1 []int, list2 []int) []int {
	interList := make([]int, 0, len(list1)) // слайс для хранения пересечения двух неупорядоченных множеств

	intersectionMap := make(map[int]struct{}) // map для проверки пересечения множеств

	// Заполняем map числами из первого множества
	for index := 0; index < len(list1); index++ {
		intersectionMap[list1[index]] = struct{}{}
	}

	// Проверяем есть ли в map число из первого множества, если есть, то значит есть пересечение
	for index := 0; index < len(list2); index++ {
		if _, ok := intersectionMap[list2[index]]; ok {
			interList = append(interList, list2[index])
		}
	}

	return interList
}

func main() {
	// Первое множество
	firstList := []int{8, 12, 5, 1, 23, 43, 84}

	// Второе множество
	secondList := []int{55, 46, 8, 19, 23, 78, 1}

	fmt.Println(Intersection(firstList, secondList))

}
