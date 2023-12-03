package main

/*
В данном коде есть две проблемы.

Во-первых, данная реализация приведет к утечке памяти, поскольку глобальная
переменная ссылается на срез, имеющий общую область памяти с бо́льшей строкой,
из-за чего сборщик мусора не сможет очистить неиспользуемую часть строки
после выхода из функции.

Во-вторых, в случае, если строка состоит из символов unicode,
операция взятия среза (v[:100]) вернет первые 100 байт, но не 100 символов.
Кроме того, после выполнения этой операции последний символ строки
может отображаться некорректно (если он занимал больше одного байта).
*/

var justString string

func someFunc() {
	v1 := []rune(createHugeString(1 << 10))
	v2 := make([]rune, 100)

	copy(v2, v1[:100]) // копируем данные в новую область памяти

	justString = string(v2)
}

func createHugeString(length int) string {
	s := make([]rune, length)
	return string(s)
}

func main() {
	someFunc()
}
