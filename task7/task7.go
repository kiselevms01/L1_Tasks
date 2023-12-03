package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcMap struct {
	mu sync.Mutex // Mьютекс для синхронизации доступа к map
	mp map[string]int
}

// Обновление данных в мапе по ключу key
func (c *ConcMap) Writer(key string) {
	c.mu.Lock()
	c.mp[key]++
	c.mu.Unlock()
}

// чтение из мапы
func (c *ConcMap) Reader(key string) int {
	return c.mp[key]
}

func main() {
	m := &ConcMap{mp: make(map[string]int)}

	key := "test"

	for i := 0; i < 100; i++ {
		go m.Writer(key)
	}

	time.Sleep(time.Second * 2)

	fmt.Printf("Получили: %d\n", m.Reader(key)) // выводим значение
}
