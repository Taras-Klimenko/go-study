package main

import (
	"fmt"
	"time"
)

// func main() {
//     for i := 1; i <= 3; i++ {
//         defer fmt.Println("defer", i)
//     }
//     fmt.Println("конец функции")
// }

//	func main() {
//	    x := 1
//	    defer fmt.Println("отложенный x:", x)
//	    x = 42
//	    fmt.Println("текущий x:", x)
//	}

func track(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s занял %v\n", name, time.Since(start))
	}
}

func main() {
	defer track("slow")() // обратите внимание на вторые скобки
	time.Sleep(600 * time.Millisecond)
}
