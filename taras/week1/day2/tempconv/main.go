package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func cToF(c float64) float64 {
	return c*9/5 + 32
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func cToK(c float64) float64 {
	return c + 273.15
}

func kToC(k float64) float64 {
	return k - 273.15
}

func main() {
	// создаём флаги
	value := flag.Float64("value", 0, "значение температуры")
	unit := flag.String("unit", "c", "исходная шкала: c или f")
	// после создания флагов парсим их и наполняем значениями
	flag.Parse()

	// flag.Float64 и flag.String возвращают УКАЗАТЕЛИ. Указатель - это адрес в памяти.
	// Разыменовывать через *value и *unit. Разыменовать - получить значение по этому адресу.

	switch strings.ToLower(*unit) {
	case "c":
		fmt.Printf("%.1f°C, %.1f°F, %.1f°K", *value, cToF(*value), cToK(*value))
	case "f":
		fmt.Printf("%.1f°C, %.1f°F, %.1f°K", fToC(*value), *value, cToK(fToC(*value)))
	case "k":
		fmt.Printf("%.1f°C, %.1f°F, %.1f°K", kToC(*value), fToC(kToC(*value)), *value)
	default:
		fmt.Printf("Ошибка: неизвестная шкала %s\n", *unit)
		os.Exit(1)

	}
}
