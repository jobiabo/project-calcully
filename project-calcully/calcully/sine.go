package calcully

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Sin(input string) (float64, error) {
	caseInput := strings.ToLower(input)
	cleanedPrefix := strings.Split(caseInput, "sin")
	removedPrefix := strings.Join(cleanedPrefix, "")
	cleanedBrace := strings.Trim(removedPrefix, "()")

	num, err := strconv.ParseFloat(cleanedBrace, 64)
	if err != nil {
		return 0, fmt.Errorf("math error")
	}
	solved := math.Sin(num)
	return (solved), nil

}
