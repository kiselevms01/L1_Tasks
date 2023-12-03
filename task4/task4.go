package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func WorkerWrite(ctx context.Context, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Write end")
			return
		default:
			ch <- rand.Intn(1000)
			time.Sleep(time.Millisecond * 1000)
		}
	}
}

func WorkerRead(ctx context.Context, ch <-chan int, w int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Read end")
			return
		default:
			fmt.Printf("Worker: %d; data: %d\n", w, <-ch)
		}
	}
}

func main() {
	var cWorker int

	fmt.Println("Количество воркеров: ")
	_, err := fmt.Scan(&cWorker)
	if err != nil {
		fmt.Printf("Ошибка ввода: %v", err)
		return
	}

	ctx, cansel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	ch := make(chan int)

	wg.Add(cWorker + 1)

	go WorkerWrite(ctx, ch, &wg)

	for i := 0; i < cWorker; i++ {
		go WorkerRead(ctx, ch, i, &wg)
	}

	// Создаём канал для обработки сигнала к завершению работы программы через Ctrl + C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT)

	// Прослушиваем канал на предмет обнаружения комбинации Ctrl + C
	<-signalChan

	// Уведомляем контексты об их завершении
	cansel()

	// Закрываем канал
	close(ch)

	// Ждем завершения всех горутин
	wg.Wait()
}
