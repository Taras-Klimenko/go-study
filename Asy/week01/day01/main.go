package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("Привет,\nGo!")
	fmt.Printf("Версия: %s\n, ядер: %d\n", runtime.Version(), runtime.NumCPU())
	file, _ := os.OpenFile("dd.txt", os.O_APPEND|os.O_WRONLY, 0404)
	fmt.Fprintln(file, 456764537368357)
	fmt.Printf("%s\n", "строка")
}
