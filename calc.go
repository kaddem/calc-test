package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var errorsMap = map[string]string{
	"rome":    "Выдача паники, так как в римской системе нет отрицательных чисел и нуля.",
	"systems": "Выдача паники, так как используются одновременно разные системы счисления.",
	"math":    "Выдача паники, так как строка не является математической операцией.",
	"format":  "Выдача паники, так как формат математической операции не удовлетворяет заданию",
}

var romanToArabicMap = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"L":    50,
	"C":    100,
}

func convertToRoman(arabic int) string {
	res := ""
	if arabic/100 == 1 {
		return "C"
	}

	if arabic >= 90 {
		res += "XC"
		arabic -= 90
	}

	if arabic >= 50 {
		res += "L"
		arabic -= 50
	}

	for arabic >= 10 {
		res += "X"
		arabic -= 10
	}

	if arabic == 9 {
		res += "IX"
		arabic -= 9
	}
	if arabic >= 5 {
		res += "V"
		arabic -= 5
	}
	if arabic == 4 {
		res += "IV"
		arabic -= 4
	}
	for arabic > 0 {
		res += "I"
		arabic -= 1
	}

	return res
}

func convertToArabic(roman string) int {
	return romanToArabicMap[roman]
}

func isRomanNum(str string) bool {
	for _, letter := range str {
		if string(letter) != "I" && string(letter) != "V" && string(letter) != "X" {
			return false
		}
	}

	return true
}

func printResult(result int, isRoman bool) {
	if isRoman && result <= 0 {
		panic(errorsMap["rome"])
	}

	if isRoman {
		fmt.Println(convertToRoman(result))
		return
	}

	fmt.Println(result)
}

func main() {
	var num1 int
	var num2 int
	var result int
	var err error
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	operands := strings.Split(input, " ")

	if len(operands) > 3 {
		panic(errorsMap["format"])
	}
	if len(operands) < 3 {
		panic(errorsMap["math"])
	}

	if isRomanNum(operands[0]) {
		num1 = convertToArabic(operands[0])
	} else {
		num1, err = strconv.Atoi(operands[0])
		if err != nil {
			panic(errorsMap["math"])
		}
	}

	if isRomanNum(operands[2]) {
		num2 = convertToArabic(operands[2])
	} else {
		num2, err = strconv.Atoi(operands[2])
		if err != nil {
			panic(errorsMap["math"])
		}
	}

	if isRomanNum(operands[0]) != isRomanNum(operands[2]) {
		panic(errorsMap["systems"])
	}

	if num1 > 10 || num2 > 10 || num1 <= 0 || num2 <= 0 {
		panic(errorsMap["format"])
	}

	switch operands[1] {
	case "-":
		result = num1 - num2
	case "+":
		result = num1 + num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		panic(errorsMap["math"])
	}

	printResult(result, isRomanNum(operands[0]))
}
