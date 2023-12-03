package main

import "fmt"

func WriteOne(arr [6]int, ch chan<- int) {
	for _, value := range arr {
		ch <- value
	}

	close(ch)
}

func WriteTwo(ch1, ch2 chan int) {
	for {
		sq, ok := <-ch1 // читаем из канала1 в канал2

		// если канал1 закрыли - закрываем канал2 и выходим из ф-и
		if !ok {
			close(ch2)
			return
		}

		ch2 <- sq * sq
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	array := [...]int{2, 6, 3, 8, 4, 10}

	go WriteOne(array, ch1)
	go WriteTwo(ch1, ch2)

	// в бесконечном цикле читаем из канала2 и выводим на экран
	for {
		result, ok := <-ch2

		// если канал2 закрыли - выходим из цикла
		if !ok {
			return
		}
		fmt.Printf("Square : %d\n", result)
	}
}
