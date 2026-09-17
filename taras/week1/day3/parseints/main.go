package main

import (
	"fmt"
	"strconv"
	"strings"
)

// parseInts разбирает строку вида "12,45,7" в срез чисел.
// Пробелы вокруг элементов игнорируются. Пустая строка даёт (nil, nil).
// При первом же нечисловом элементе возвращает (nil, ошибку).
func parseInts(s string) ([]int, error) {
	// пустая строка возвращает nil, nil
	if strings.TrimSpace(s) == "" {
		return nil, nil
	}
	stringArr := strings.Split(s, ",")

	resultSlice := []int{}

	for i := 0; i < len(stringArr); i++ {
		result, error := strconv.Atoi(strings.TrimSpace(stringArr[i]))
		if error != nil {
			return nil, error
		}
		resultSlice = append(resultSlice, result)

	}
	return resultSlice, nil
}

func main() {
	cases := []string{
		"12,45,7",
		"12, 45 , 7",
		"",
		"12,abc,7",
		"12,,7",
	}

	for _, c := range cases {
		nums, err := parseInts(c)
		if err != nil {
			fmt.Printf("%-15q → ошибка: %v\n", c, err)
			continue
		}
		fmt.Printf("%-15q → %v\n", c, nums)
	}
}
