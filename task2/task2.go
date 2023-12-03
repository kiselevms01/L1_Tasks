package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	array := [...]int{2, 4, 6, 8, 10} // Входной массив

	arrSquare := [5]int{}

	for i, val := range array {
		wg.Add(1) // В цикле для каждой итерации инкрементируем счетчик горутин
		go func(i int, val int) {
			arrSquare[i] = val * val
			wg.Done() // в конце дикрементируем счетчик
		}(i, val)
	}

	wg.Wait() // Ожидаем окончание выполнения всех горутин

	// и выводим результат на экран в основной горутине:
	fmt.Printf("Квадраты элементов массива %v : %v", array, arrSquare)
}
