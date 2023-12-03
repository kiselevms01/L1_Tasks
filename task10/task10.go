package main

import (
	"fmt"
)

// Функция объединяет температуры в группы с шагом в 10 градусов
func TempSort(t []float64) map[int][]float64 {
	tSort := make(map[int][]float64) // создадим мапу слайсов

	for _, val := range t {
		key := (int(val / 10)) * 10          // получаем значение ключа
		tSort[key] = append(tSort[key], val) // кладем в слайс по ключу соответсвующее значение
	}

	return tSort
}

func PrintTemp(mp map[int][]float64) {
	for k, slice := range mp {
		fmt.Printf("%d : %v\n", k, slice)
	}

}

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	PrintTemp(TempSort(temp))
}
