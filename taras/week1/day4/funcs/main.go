package main

import "fmt"

// Map возвращает НОВЫЙ срез с результатами применения f к каждому элементу.
// Исходный срез не изменяется.
func Map(nums []int, f func(int) int) []int {
	result := make([]int, len(nums)) // делаем слайс из такого же количества элементов, как в исходном nums
	for i, value := range nums {
		result[i] = f(value)
	}
	return result
}

// Filter возвращает НОВЫЙ срез из элементов, для которых pred вернул true.
func Filter(nums []int, pred func(int) bool) []int {
	result := make([]int, 0, len(nums)) // делаем пустой слайс с емкостью как у исходного массива

	for _, value := range nums {
		if pred(value) == true { // можно просто if pred(value)
			result = append(result, value)
		}
	}
	return result
}

// Sum возвращает сумму элементов. Для пустого среза — 0.
func Sum(nums []int) int {
	result := 0
	for _, value := range nums {
		result = result + value // можно просто result += value
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	squares := Map(evens, func(n int) int { return n * n })

	fmt.Println("исходный:", nums)
	fmt.Println("чётные:  ", evens)
	fmt.Println("квадраты:", squares)
	fmt.Println("сумма:   ", Sum(squares))
}
