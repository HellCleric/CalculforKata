package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanMap = map[string]int{
	"I": 1, "V": 5, "X": 10, "XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

var arabicMap = map[int]string{
	1: "I", 5: "V", 10: "X", 40: "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func romanToArabicNumeral(roman string) int {
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

var arabic bool

func arabicToRomanNumeral(num int) string {
	result := ""

	for num > 0 {
		for key := 1000; key >= 1; key /= 10 {
			value := arabicMap[key]
			if num >= key {
				result += value
				num -= key
			}
		}
	}

	return result
}

func calcul(a, b string, operator string, isRoman bool) string {
	var numA, numB int

	if !isRoman {
		numA, _ = strconv.Atoi(a)
		numB, _ = strconv.Atoi(b)
	} else {
		numA = romanToArabicNumeral(a)
		numB = romanToArabicNumeral(b)
	}

	switch operator {
	case "+":
		result := numA + numB
		if isRoman {
			return arabicToRomanNumeral(result)
		}
		return strconv.Itoa(result)
	case "-":
		result := numA - numB
		if isRoman {
			return arabicToRomanNumeral(result)
		}
		return strconv.Itoa(result)
	case "*":
		result := numA * numB
		if isRoman {
			return arabicToRomanNumeral(result)
		}
		return strconv.Itoa(result)
	case "/":
		if numB == 0 {
			panic("Деление на ноль")
		}
		result := numA / numB
		if isRoman {
			return arabicToRomanNumeral(result)
		}
		return strconv.Itoa(result)
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
			panic("Неправильный формат выражения")
		}

		a := parts[0]
		operator := parts[1]
		b := parts[2]
		result := calcul(a, b, operator, false)
		if _, err := strconv.Atoi(a); err != nil {
			result = calcul(a, b, operator, true)
		}
		for _, c := range a {
			if c == '0' {
				arabic = false
				panic("Нельзя использовать 0")
			}
		}
		for _, c := range b {
			if c == '0' {
				arabic = false
				panic("нельзя использовать 0")
			}

		}

		if arabic {
			numA, _ := strconv.Atoi(a)
			numB, _ := strconv.Atoi(b)
			if numA < 1 || numA > 10 || numB < 1 || numB > 10 {
				panic("Числа должны быть от 1 до 10")
			}
			result = calcul(a, b, operator, false)
		} else {
			nA := romanToArabicNumeral(a)
			nB := romanToArabicNumeral(b)

			if nA < romanMap["I"] || nA > romanMap["X"] || nB < romanMap["I"] || nB > romanMap["X"] {
				panic("Числа должны быть от I до X")
			}

			result = calcul(a, b, operator, true)
			if romanToArabicNumeral(result) < 1 {
				panic("Результат не может быть меньше единицы")
			}

		}

		fmt.Println("Результат:", result)
	}
}
