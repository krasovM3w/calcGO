package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
var arabicToRoman = map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV", 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX",
	21: "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV", 26: "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX", 30: "XXX",
	31: "XXXI", 32: "XXXII", 33: "XXXIII", 34: "XXXIV", 35: "XXXV", 36: "XXXVI", 37: "XXXVII", 38: "XXXVIII", 39: "XXXIX", 40: "XL",
	41: "XLI", 42: "XLII", 43: "XLIII", 44: "XLIV", 45: "XLV", 46: "XLVI", 47: "XLVII", 48: "XLVIII", 49: "XLIX", 50: "L",
	51: "LI", 52: "LII", 53: "LIII", 54: "LIV", 55: "LV", 56: "LVI", 57: "LVII", 58: "LVIII", 59: "LIX", 60: "LX",
	61: "LXI", 62: "LXII", 63: "LXIII", 64: "LXIV", 65: "LXV", 66: "LXVI", 67: "LXVII", 68: "LXVIII", 69: "LXIX", 70: "LXX",
	71: "LXXI", 72: "LXXII", 73: "LXXIII", 74: "LXXIV", 75: "LXXV", 76: "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX", 80: "LXXX",
	81: "LXXXI", 82: "LXXXII", 83: "LXXXIII", 84: "LXXXIV", 85: "LXXXV", 86: "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX", 90: "XC",
	91: "XCI", 92: "XCII", 93: "XCIII", 94: "XCIV", 95: "XCV", 96: "XCVI", 97: "XCVII", 98: "XCVIII", 99: "XCIX"}

func RomanToArabic(roman string) (int, error) {
	value, err := romanToArabic[roman]
	if !err {
		return 0, fmt.Errorf("Ошибка ввода: %s ", roman)
	}

	return value, nil
}

func ArabicToRoman(arabic int) (string, error) {
	value, err := arabicToRoman[arabic]
	if !err {
		return "0", fmt.Errorf("Ошибка ввода : %s ", arabic)
	}

	return value, nil
}

func main() {

	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("введите выражение: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Удаление символа новой строки и пробельных символов с обеих сторон

		// Регулярное выражение с учетом пробелов вокруг оператора
		re := regexp.MustCompile(`^(\d+|[IVX]+)\s*([\+\-\*/])\s*(\d+|[IVX]+)$`)
		matches := re.FindStringSubmatch(input)
		if matches == nil {
			fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
			return
		}
		//fmt.Println("вывод:", matches)

		var a, b int
		var _ error

		if _, err := strconv.Atoi(matches[1]); err == nil {
			a, _ = strconv.Atoi(matches[1])
			b, err = strconv.Atoi(matches[3])
			if err != nil {
				fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления")
			}

		} else {
			a, err = RomanToArabic(matches[1])
			if err != nil {
				fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
				return
			}
			b, err = RomanToArabic(matches[3])
			if err != nil {
				fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
				return
			}
		}

		if a < 1 || a > 10 || b < 1 || b > 10 {
			fmt.Println("Введите число более 1 или менее 10")
		}
		var result int
		switch matches[2] {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			result = a / b
			if b == 0 {
				fmt.Println("на ноль делить нельзя!")
			}
		default:
			fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
		}
		//fmt.Println(a, b)
		if _, err := strconv.Atoi(matches[1]); err == nil {
			fmt.Println(result)
		} else {
			if result < 1 {
				fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
				return
			}
			romanResult, err := ArabicToRoman(result)
			if err != nil {
				fmt.Println("Ошибка ввода: ", err)
				return
			}
			fmt.Println(romanResult)
		}
	}
}

// с пивом пойдет
