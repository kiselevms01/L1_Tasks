package main

import (
	"fmt"
)

func main() {
	var k, number int

	fmt.Println("Введите количество элементов:")
	fmt.Scan(&k)

	var array = make([]int, k)

	// Созданим отсортированный массив
	for i := 0; i < k; i++ {
		array[i] = i - k/2
	}
	fmt.Printf("Массив: %v\n", array)

	fmt.Println("Введите искомое число:")
	fmt.Scan(&number)

	var binId, ok = binarySearch(array, number)
	if ok {
		fmt.Printf("Число %d найдено по индексу %d", array[binId], binId)
	} else {
		fmt.Println("Число не найдено!")
	}
}

// Функция бинаргного поиска
func binarySearch(array []int, lookingFor int) (int, bool) {
	var mid, find int

	// устанавливаем границы поиска
	left := 0
	right := len(array) - 1

	for left <= right {
		mid = (left + right) / 2
		find = array[mid]
		if find == lookingFor { // сравниваем среднее значение с искомым, если равны то возвращаем индекс
			return mid, true
		}
		if find > lookingFor {
			right = mid - 1 // если среднее значение > искомого, то правая граница обновляется
		} else {
			left = mid + 1 // если среднее значение < искомого, то обновляется левая граница
		}
	}
	return 0, false // возвращаем false в случае безуспешного поиска
}
