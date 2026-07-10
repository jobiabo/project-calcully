package calcully

import (
	"fmt"
	"log"
	"strconv"
)

func Solve(sTringInput string, number []string) (float64, error) {
	log.Println("confirming stringinput:", sTringInput)
	log.Println("confirming number:", number)

	result := 0.00
	newToken := []float64{}

	for _, tok := range number {
		num, err := strconv.ParseFloat(tok, 64)
		if err != nil {
			return result, fmt.Errorf("math error")
		}
		newToken = append(newToken, num)
	}

	j := 0
	for i := 0; i < len(sTringInput); i++ {
		if i == 0 && j < len(newToken) {
			result += newToken[j]
			j++
		}

		// Fixed: Read characters properly and fixed index logic
		switch sTringInput[i] {
		case '+':
			if j < len(newToken) {
				result += newToken[j]
				j++
			}
		case '-':
			if j < len(newToken) {
				result -= newToken[j]
				j++
			}

		case '*':
			if j < len(newToken) {
				result *= newToken[j]
				j++
			}

		case '/':
			if j < len(newToken) {
				if newToken[j] == 0 {
					return 0, fmt.Errorf("math syntax error")
				}
				result /= newToken[j]
				j++
			}
		}
	}
	return result, nil
}
