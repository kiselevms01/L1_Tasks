package main

import (
	"fmt"
	"math/big"
)

// сложение
func Addition(a *big.Int, b *big.Int) *big.Int {
	add := new(big.Int).Add(a, b)
	return add
}

// вычитание
func Subtraction(a *big.Int, b *big.Int) *big.Int {
	sub := new(big.Int).Sub(a, b)
	return sub
}

// умножение
func Multiplication(a *big.Int, b *big.Int) *big.Int {
	mul := new(big.Int).Mul(a, b)
	return mul
}

// деление
func Division(a *big.Int, b *big.Int) *big.Int {
	div := new(big.Int).Div(a, b)
	return div
}

func main() {

	// Создаём переменных типа big.Int
	a := big.NewInt(0)
	b := big.NewInt(0)

	// Устанавливаем значения в переменные a и b через SetString, указывая само значение и тип системы счисления
	a.SetString("978697859596978878596586878595678", 10)
	b.SetString("223342515343435254445521543352434", 10)

	fmt.Println(Addition(a, b))       // Сложение
	fmt.Println(Subtraction(a, b))    // Вычитание
	fmt.Println(Multiplication(a, b)) // Умножение
	fmt.Println(Division(a, b))       // Деление

}
