package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestNumberToThaiText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "integer only",
			input:    "1234",
			expected: "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน",
		},
		{
			name:     "with satang",
			input:    "33333.75",
			expected: "สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์",
		},
		{
			name:     "zero",
			input:    "0",
			expected: "ศูนย์บาทถ้วน",
		},
		{
			name:     "twenty one",
			input:    "21",
			expected: "ยี่สิบเอ็ดบาทถ้วน",
		},
		{
			name:     "one hundred one",
			input:    "101",
			expected: "หนึ่งร้อยเอ็ดบาทถ้วน",
		},

		// ---------- ล้าน ----------
		{
			name:     "one million",
			input:    "1000000",
			expected: "หนึ่งล้านบาทถ้วน",
		},
		{
			name:     "ten million",
			input:    "10000000",
			expected: "สิบล้านบาทถ้วน",
		},
		{
			name:     "one hundred million",
			input:    "100000000",
			expected: "หนึ่งร้อยล้านบาทถ้วน",
		},

		// ---------- พันล้าน ----------
		{
			name:     "one billion",
			input:    "1000000000",
			expected: "หนึ่งพันล้านบาทถ้วน",
		},
		{
			name:     "two billion one hundred million",
			input:    "2100000000",
			expected: "สองพันหนึ่งร้อยล้านบาทถ้วน",
		},
		{
			name:     "billion with satang",
			input:    "1234567890.50",
			expected: "หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทห้าสิบสตางค์",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := decimal.RequireFromString(tt.input)
			result := numberToThaiText(d)

			if result != tt.expected {
				t.Errorf("input %s: expected %q, got %q",
					tt.input, tt.expected, result)
			}
		})
	}
}
