package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct {
	a, b int
	op   string
}

func (calc Calc) calculate() int {
	if calc.op == "/" {
		return calc.a / calc.b
	}
	if calc.op == "*" {
		return calc.a * calc.b
	}
	if calc.op == "+" {
		return calc.a + calc.b
	}
	if calc.op == "-" {
		return calc.a - calc.b
	}
	panic("Wrong operation")
}

func romanToArabic(num string) int {
	translateRoman := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100}
	var decNum, tmpNum int
	for i := len(num) - 1; i >= 0; i-- {
		romanDigit := num[i]
		decDigit := translateRoman[romanDigit]
		if decDigit < tmpNum {
			decNum -= decDigit
		} else {
			decNum += decDigit
			tmpNum = decDigit
		}
	}
	if decNum < 1 || decNum > 10 {
		panic("Number not in range")
	}
	return decNum
}

func arabicToRoman(num int) string {
	if num < 1 || num > 100 {
		panic("Resul not in range")
	}
	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for num >= conversion.value {
			roman.WriteString(conversion.digit)
			num -= conversion.value
		}
	}
	return roman.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input string to calculation")
	text, _ := reader.ReadString('\n')
	arr := strings.Fields(text)
	if len(arr) != 3 {
		panic("Incorrect input format")
	}
	a, errA := strconv.Atoi(arr[0])
	b, errB := strconv.Atoi(arr[2])
	op := arr[1]
	if errA == nil && errB == nil {
		calc := Calc{a, b, op}
		fmt.Println(calc.calculate())
		return
	}
	if errA != nil && errB != nil {
		a := romanToArabic(arr[0])
		b := romanToArabic(arr[2])
		if a <= b && op == "-" || a < b && op == "/" {
			panic("Result must be in range [I; C]")
		}
		calc := Calc{a, b, op}
		fmt.Println(arabicToRoman(calc.calculate()))
		return
	}
	panic("All numbers must be arabic or roman")
}
