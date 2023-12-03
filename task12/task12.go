package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(GetSet(arr))
}

func GetSet(data []string) []string {

	maps := make(map[string]struct{})
	for _, v := range data {
		maps[v] = struct{}{}
	}

	// создадим слайс и добавим в него ключи из мапы (уникальные элементы)
	result := make([]string, 0, len(maps))
	for k := range maps {
		result = append(result, k)
	}
	return result
}
