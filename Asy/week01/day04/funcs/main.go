package main

import (
	"fmt"
)

// Map возвращает НОВЫЙ срез с результатами применения f к каждому элементу.
// Исходный срез не изменяется.
func Map(nums []int, f func(int) int) []int {
	if nums == nil {
		return nil
	}
	res := make([]int, len(nums))
	for i, s := range nums {
		res[i] = f(s)
		// fmt.Print("первая",res)
	}

	return res

}

// Filter возвращает НОВЫЙ срез из элементов, для которых pred вернул true.
func Filter(nums []int, pred func(int) bool) []int {
	if nums == nil {
		return nil
	}
	res := make([]int, 0, len(nums))
	for _, s := range nums {
		if pred(s) {
			res = append(res, s)
			// fmt.Print("вторая",res)
		}

	}
	return res
	// ваш код
}

// Sum возвращает сумму элементов. Для пустого среза — 0.
func Sum(nums []int) int {
	// if nums == nil {
	// 	return 0
	// }
	res := 0
	for _, s := range nums {
		res += s

	}
	return res
	// ваш код
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
