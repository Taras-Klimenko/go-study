package main

import (
	"fmt"
)

func appendTest() {
	var s []int
	prev := -1
	for i := 0; i < 20; i++ {
		s = append(s, i)
		if cap(s) != prev {
			fmt.Printf("len=%2d cap=%2d\n", len(s), cap(s))
			prev = cap(s)
		}
	}
}
func commonSliceTest() {
	a := []int{1, 2, 3, 4, 5}
	b := a[1:3]
	b[0] = 99
	fmt.Println(a, b)
}

func capacityTest() {
	a := []int{1, 2, 3, 4, 5}
	b := a[1:3]
	fmt.Println(len(b), cap(b)) // сначала посмотрите на cap

	b = append(b, 777)
	fmt.Println(a, b)
}

func main() {
	// Способ 1: полная копия
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 2)
	copy(b, a[1:3])
	b = append(b, 777)
	fmt.Println(a, b) // a цел

	// Способ 2: трёхиндексный срез — обрезаем cap
	c := []int{1, 2, 3, 4, 5}
	d := c[1:3:3]
	fmt.Println(len(d), cap(d))
	d = append(d, 777)
	fmt.Println(c, d) // c цел
}
