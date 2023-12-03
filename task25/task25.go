package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	Sleep1(3 * time.Second)
	fmt.Println(time.Now().Sub(now))

	now = time.Now()
	Sleep3(3 * time.Second)
	fmt.Println(time.Now().Sub(now))
}

func Sleep1(tm time.Duration) {
	<-time.After(tm) // ждем истечения продолжительности tm и отпр. текущее время по каналу
}

// с использованием контекста
func Sleep3(tm time.Duration) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), tm)
	defer cancelFunc()
	<-timeout.Done()
}
