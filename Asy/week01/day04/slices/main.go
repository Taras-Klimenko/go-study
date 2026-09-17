package main

import (
	"fmt"
)

// func main() {
// 	var s []int
// 	prev := -1
// 	for i := 0; i < 20; i++ {
// 		s = append(s, i)
// 		if cap(s) != prev {
// 			fmt.Printf("len=%2d cap=%2d\n", len(s), cap(s))
// 			prev = cap(s)
// 		}
// 	}
// }

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	b := a[1:3]
// 	b[0] = 99
// 	fmt.Println(a, b) //a{99,2,3,4,5},b{99,3}
// }

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	b := a[1:3]
// 	fmt.Println(len(b), cap(b)) // сначала посмотрите на cap
// 	//len = 2, cap =4?
// 	b = append(b, 777)
// 	fmt.Println(a, b)
// }

// a := []int{1, 2, 3, 4, 5}
// b := make([]int, 2)
// copy(b, a[1:3])копируем{2,3}
// b = append(b, 777) {2,3,777}
// fmt.Println(a, b)   // a цел

// // Способ 2: трёхиндексный срез — обрезаем cap
// c := []int{1, 2, 3, 4, 5}
// d := c[1:3:3]
// fmt.Println(len(d), cap(d))
// d = append(d, 777)
// fmt.Println(c, d)
