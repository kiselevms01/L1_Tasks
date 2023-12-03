package main

import "fmt"

// Встраиваемая структура
type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human // встроенное (или анонимное) поле
	Name  string
}

// определим 2 простых метода для структуры Human:
func (h *Human) PrintName() {
	fmt.Printf("Имя водителя: %s\n", h.Name)
}

func (h *Human) PrintAge() {
	fmt.Printf("Его возраст: %d\n", h.Age)
}

func Embedding() {

	newAction := Action{
		Human{"Василий", 30},
		"В пути",
	}

	fmt.Printf("Данные: %s, %d лет. Действие: %s\n", newAction.Human.Name, newAction.Age, newAction.Name)

	// вызовем методы структуры Human для переменной типа Action:
	newAction.PrintName()
	newAction.PrintAge()
}

func main() {
	fmt.Printf("Result:\n")
	Embedding()
}
