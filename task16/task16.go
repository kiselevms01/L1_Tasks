package main

import (
	"fmt"
	"math/rand"
)

func sort(arr []int, start, end int) {
	if start >= end {
		return
	}

	// Выбираем опорный элемент, который будет использоваться для разделения среза на две части
	pivot := arr[start]

	i := start + 1

	for j := start; j <= end; j++ {

		// Если текущий элемент меньше опорного элемента, меняем его местами с элементом на позиции i и увеличиваем i
		if pivot > arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// Помещаем опорный элемент на правильную позицию, чтобы все элементы слева от него были меньше, а справа - больше
	arr[start], arr[i-1] = arr[i-1], arr[start]

	// Рекурсивно вызываем функцию sort для левой части среза (элементы меньше опорного элемента)
	sort(arr, start, i-2)

	// Рекурсивно вызываем функцию sort для правой части среза (элементы больше опорного элемента)
	sort(arr, i, end)
}

func main() {
	array := make([]int, 100)

	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(len(array))
	}

	fmt.Println("Source list:", array)
	sort(array, 0, len(array)-1)
	fmt.Println("Sorted list:", array)
}
