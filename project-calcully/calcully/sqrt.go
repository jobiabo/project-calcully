package calcully

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Sqrt(input string) (float64, error) {
	caseInput := strings.ToLower(input)
	cleanedPrefix := strings.Split(caseInput, "sqrt")
	removedPrefix := strings.Join(cleanedPrefix, "")
	cleanedBrace := strings.Trim(removedPrefix, "()")

	num, err := strconv.ParseFloat(cleanedBrace, 64)
	if err != nil {
		return 0, fmt.Errorf("math error")
	}
	solved := math.Sqrt(num)
	return (solved), nil

}
