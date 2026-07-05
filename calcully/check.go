package calcully

import (
	"log"
	"strings"
)

func Check(s string) ([]string, string) {
	cleeanedA := strings.ReplaceAll(s, "÷", "/")
	cleeanedB := strings.ReplaceAll(cleeanedA, "×", "*")
	log.Println("confirming input:", cleeanedB)
	ops := "+-*/" // Fixed: Added quotes around operators
	return strings.FieldsFunc(cleeanedB, func(r rune) bool {
		return strings.ContainsRune(ops, r)
	}), cleeanedB
}
