package calcully

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Cos(input string) (float64, error) {
	caseInput := strings.ToLower(input)
	cleanedPrefix := strings.Split(caseInput, "cos")
	removedPrefix := strings.Join(cleanedPrefix, "")
	cleanedBrace := strings.Trim(removedPrefix, "()")

	num, err := strconv.ParseFloat(cleanedBrace, 64)
	if err != nil {
		return 0, fmt.Errorf("math error")
	}
	solved := math.Cos(num)
	return (solved), nil

}
