package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция записи данных в канал в бесконечном цикле
func Write(ch chan int) {
	for {
		time.Sleep(time.Millisecond * 500)
		ch <- rand.Intn(1000)
	}
}

func main() {
	var n int

	fmt.Println("Введите количество секунд: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Printf("Ошибка ввода: %v", err)
		return
	}

	ch := make(chan int)
	now := time.Now()

	go Write(ch)

	// Итерируемся по каналу и читаем данные в основной горутине
	for i := range ch {

		fmt.Printf("Новое значение из канала: %d\n", i)

		// если время вышло, выводим сообщение и завершаем программу
		if time.Now().Sub(now) >= time.Duration(n)*time.Second {
			fmt.Printf("End: Время (%dsec) вышло.\n", n)
			close(ch) // или же return
			fmt.Println("Канал зыкрыт")
		}
	}
}
