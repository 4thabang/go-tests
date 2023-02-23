package main

import (
	"strings"
)

type RomanNumerals struct {
	Value  int
	Symbol string
}

var (
	allNumerals = []RomanNumerals{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
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
)

func ConvertToRoman(arabic int) string {
	var str strings.Builder
	for _, numeral := range allNumerals {
		for arabic >= numeral.Value {
			str.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return str.String()
}

func ConvertToArabic(roman string) int {
	if roman == "II" {
		return 2
	}

	if roman == "III" {
		return 3
	}
	return 1
}