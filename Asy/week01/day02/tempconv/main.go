package main

// - [ ] Формулы: `F = C×9/5 + 32`, `C = (F−32)×5/9`
// - [ ] Вывод форматировать через `fmt.Printf` с `%.1f`

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func cToF(c float64) float64 {
	return c*9/5 + 32
	// ваш код
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
	// ваш код
}

func main() {
	value := flag.Float64("value", -40, "значение температуры")
	unit := flag.String("unit", "c", "исходная шкала: c или f")
	flag.Parse()

	// flag.Float64 и flag.String возвращают УКАЗАТЕЛИ.
	// Разыменовывать через *value и *unit.

	switch strings.ToLower(*unit) {
	case "c":
		c := cToF(*value)
		fmt.Printf("%.1fC= %.1fF", *value, c)
	case "f":
		f := fToC(*value)
		fmt.Printf("%.1fF = %.1fC", *value, f)
	default:
		fmt.Printf("ERROR")
		os.Exit(1)
		// ваш код: две ветки и default
		// в default — сообщение об ошибке и os.Exit(1)
	}
}
