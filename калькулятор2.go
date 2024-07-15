package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToArabicNumeral(roman string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
	}

	result := 0
	prev := 0

	for i := len(roman) - 1; i >= 0; i-- {
		current := romanMap[string(roman[i])]

		if current >= prev {
			result += current
		} else {
			result -= current
		}

		prev = current
	}

	return result
}

func calcul(a, b string, operator string) int {
	numA := romanToArabicNumeral(a)
	numB := romanToArabicNumeral(b)

	switch operator {
	case "+":
		return numA + numB
	case "-":
		return numA - numB
	case "*":
		return numA * numB
	case "/":
		return numA / numB
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
		var result int
		arabic := true
		for _, c := range a {
			if c == '0' {
				arabic = false
				panic("Нельзя использовать 0")
			}
		}
		for _, c := range b {
			if c == '0' {
				arabic = false
				panic("нельзя искользовать 0")
			}
		}

		if arabic {
			numA, _ := strconv.Atoi(a)
			numB, _ := strconv.Atoi(b)
			if numA < 0 || numA > 10 || numB < 0 || numB > 10 {
				panic("Числа должны быть от 1 до 10")
			}
			result = calcul(a, b, operator)
		} else {
			nA := romanToArabicNumeral(a)
			nB := romanToArabicNumeral(b)

			if nA < 'I' || nA > 'X' || nB < 'I' || nB > 'X' {
				panic("Числа должны быть от I до X")

			}

			if nA < 'I' || nA > 'X' || nB < 'I' || nB > 'X' {
				panic("Числа должны быть от I до X")

			}
			result = calcul(a, b, operator)
			if result < 1 {
				panic("Результат не может быть меньше единицы")
			}
		}

		fmt.Println("Результат:", result)
	}
}
