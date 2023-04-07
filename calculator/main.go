package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner((os.Stdin))
	for {
		fmt.Print("Введите выражение: ")
		scanner.Scan()
		expression := scanner.Text()

		if expression == "" {
			fmt.Println("Выход")
			return
		}

		Calculator(expression)
	}

}

func Calculator(s string) {
	// чистим пробелы в начале и в конце// на всякий случай
	s = strings.TrimSpace(s)
	// извлекаем числа и знак операции
	var a, b int
	var op string
	if i := strings.IndexAny(s, "+-*/"); i != -1 { // проверяем наличие знака операции
		op = string(s[i])
		a, _ = strconv.Atoi(strings.TrimSpace(s[:i]))   // извлекаем первое число
		b, _ = strconv.Atoi(strings.TrimSpace(s[i+1:])) // извлекаем второе число
	} else {
		fmt.Println("Недопустимый знак операции")
		return
	}

	// выполнение математической операции в зависимости от знака операции
	switch op {
	case "+":
		fmt.Print(a, " + ", b, " = ")
		fmt.Println(a + b)
	case "-":
		fmt.Print(a, " - ", b, " = ")
		fmt.Println(a - b)
	case "*":
		fmt.Print(a, " * ", b, " = ")
		fmt.Println(a * b)
	case "/":
		fmt.Print(a, " / ", b, " = ")
		fmt.Println(float64(a) / float64(b))
	default:
		fmt.Println("Недопустимый знак операции")
	}
}
