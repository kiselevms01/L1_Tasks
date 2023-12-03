package main

import (
	"fmt"
)

// Функция проверяет наличие установленного или неустановленного бита
func hasBit(num int64, pos uint) bool {
	val := num & (1 << pos)
	return val > 0
}

// Функция изменяет бит N с 1 на 0 и наоборот
func ChangeNBit(num int64, pos uint) int64 {
	pos-- // изменяет 1-8 на 0-7

	if hasBit(num, pos) {
		num &= ^(1 << pos)
	} else {
		num |= (1 << pos)
	}
	return num
}

func main() {
	num := int64(13)
	fmt.Printf("Number: %d\n in bytes: %08b\n", num, num)

	// меняем биты на позициях 1-8
	changedNum := ChangeNBit(num, 3)
	fmt.Printf("Number: %d\n in bytes:  %08b\n", changedNum, changedNum)
}
