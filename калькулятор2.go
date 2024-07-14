package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func roman(r string) int {
	romanDict := map[rune]int{'I': 1,'V': 5,'X': 10}
	result := 0
	prev := 0
	for _, c := range r {
		value, ok := romanDict[c]
		if !ok {
			return -1
		}
		if value > prev {
			result += value - 2*prev
		} else {
			result += value
		}
		prev = value
	}
	return result
}

func calcul(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("Неправильный оператор")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите выражение: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Неправильный формат выражения")
			continue
		}

		a := parts[0]
		operator := parts[1]
		b := parts[2]

		arabic := true
		for _, c := range a {
			if c < '1' || c > '9' {
				arabic = false
				panic("Числа должны быть от 1 до 10")
				break
			}
		}
		for _, c := range b {
			if c < '1' || c > '9' {
				arabic = false
				panic("Числа должны быть от 1 до 10")
				break
			}
		}

		var result int
		if arabic {
			numA, _ := strconv.Atoi(a)
			numB, _ := strconv.Atoi(b)
			result = calcul(numA, numB, operator)
		} else {
			numA := roman(a)
			numB := roman(b)
			if numA < 1 || numA > 10 || numB < 1 || numB > 10 {
				panic("Числа должны быть от I до X")
			}
			result = calcul(numA, numB, operator)
			if result < 1 {
				panic("Результат не может быть меньше единицы")
			}
		}

		fmt.Println("Результат:", result)
	}
}
