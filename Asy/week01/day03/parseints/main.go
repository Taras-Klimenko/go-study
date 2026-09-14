package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInts(s string) ([]int, error) {

	if strings.TrimSpace(s) == "" {
		return nil, nil
	}
	res := strings.Split(s, ",")
	resSlice := []int{}

	for _, r := range res {
		nums, err := strconv.Atoi(strings.TrimSpace(r))
		if err != nil {

			return nil, err
		}
		resSlice = append(resSlice, nums)
	}

	return resSlice, nil
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
