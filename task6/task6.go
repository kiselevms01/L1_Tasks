package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func OnContext(ctx context.Context, wg *sync.WaitGroup, msg string) {
	defer wg.Done()
	<-ctx.Done()
	fmt.Println("\rClosed", msg)
}

func OnChannel(s <-chan struct{}, wg *sync.WaitGroup, msg string) {
	defer wg.Done()
	<-s
	fmt.Println("\rClosed", msg)
}

func Cancel() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ctx, cancelFunc := context.WithCancel(context.Background())
	go OnContext(ctx, &wg, "1 - WithCancel")
	time.Sleep(time.Second)
	cancelFunc()
	wg.Wait()
}

func Timeout() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ctxTime, cancelFuncTime := context.WithTimeout(context.Background(), time.Second)
	defer cancelFuncTime()
	go OnContext(ctxTime, &wg, "2 - WithTimeout")
	<-ctxTime.Done()
	wg.Wait()
}

func Deadline() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	deadline, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancelFunc()
	go OnContext(deadline, &wg, "3 - WithDeadline")
	<-deadline.Done()
	wg.Wait()
}

func CtxSignal() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	fmt.Println("\tpress CTRL+C to close NotifyContext goroutine")
	deadline, cancelFunc := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancelFunc()
	go OnContext(deadline, &wg, "4 with context signal CTRL+C - NotifyContext")
	<-deadline.Done()
	wg.Wait()
}

func Channel() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	cn := make(chan struct{})
	go OnChannel(cn, &wg, "5 - WithChannel")
	time.Sleep(time.Second)
	cn <- struct{}{}
	wg.Wait()
}

func main() {
	// С помощью функции cancel() который получили с контекста, мы имеем возможность сообщить горутинам о прекращении работы.
	Cancel()
	// Контекст в котором через заданное время будет послан сигнал о прекращении работы горутинам
	Timeout()
	// То же самое что и выше - но работать будет до наступления времени дедлайна
	Deadline()
	// Тоже контекст, но привязанный к сигналам, а данном случае (ctrl + c)
	CtxSignal()
	// Альтернативные способы с помощью каналов ...
	Channel()
}
