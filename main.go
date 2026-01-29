package main

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

var thaiDigits = []string{
	"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่",
	"ห้า", "หก", "เจ็ด", "แปด", "เก้า",
}

var thaiUnits = []string{
	"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน",
}

func moneyToThaiText(num decimal.Decimal) string {
	num = num.Round(2)

	parts := strings.Split(num.StringFixed(2), ".")
	bahtPart := parts[0]
	satangPart := parts[1]

	bahtText := readNumber(bahtPart) + "บาท"

	if satangPart == "00" {
		return bahtText + "ถ้วน"
	}

	satangText := readNumber(satangPart) + "สตางค์"
	return bahtText + satangText
}

func readNumber(numStr string) string {
	if numStr == "0" {
		return thaiDigits[0]
	}

	if len(numStr) > 6 {
		millionPos := len(numStr) - 6
		millionPart := numStr[:millionPos]
		remainingPart := numStr[millionPos:]

		millionText := readNumber(millionPart)
		remainingText := readNumber(remainingPart)

		if remainingText == "" || remainingText == "ศูนย์" {
			return millionText + "ล้าน"
		}
		return millionText + "ล้าน" + remainingText
	}

	result := ""
	length := len(numStr)

	for i, ch := range numStr {
		digit := int(ch - '0')
		pos := length - i - 1

		if digit == 0 {
			continue
		}

		if pos == 0 && digit == 1 && length > 1 {
			result += "เอ็ด"
		} else if pos == 1 && digit == 2 {
			result += "ยี่"
		} else if pos == 1 && digit == 1 {
			result += ""
		} else {
			result += thaiDigits[digit]
		}

		if pos < len(thaiUnits) {
			result += thaiUnits[pos]
		}
	}

	return result
}

func main() {
	tests := []string{
		"1234",
		"33333.75",
		"21",
		"101.50",
	}

	for _, t := range tests {
		d := decimal.RequireFromString(t)
		fmt.Printf("%s -> %s\n", t, moneyToThaiText(d))
	}
}
