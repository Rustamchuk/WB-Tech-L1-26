package main

import "fmt"

/*
Дана последовательность температурных колебаний:
-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.


Пример:
-20:{-25.0, -27.0, -21.0},
10:{13.0, 19.0, 15.5},
20: {24.5}, etc.
*/

func main() {
	var (
		nums    = []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
		tempMap = make(map[int][]float32)
	)

	for _, num := range nums {
		// Ключем будем количество десятков в текущем числе * 10.
		// / 10 с округлением вниз, получаем десятки. Умножаем на 10, получается необходимый шаг в 10 градусов
		key := int(num) / 10 * 10
		// Кладем в мапу, если раньше не было такого ключа, то массив создастся сам
		tempMap[key] = append(tempMap[key], num)
	}
	fmt.Println(tempMap)
}
