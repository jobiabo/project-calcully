package calcully

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Tan(input string) (float64, error) {
	caseInput := strings.ToLower(input)
	cleanedPrefix := strings.Split(caseInput, "tan")
	removedPrefix := strings.Join(cleanedPrefix, "")
	cleanedBrace := strings.Trim(removedPrefix, "()")

	num, err := strconv.ParseFloat(cleanedBrace, 64)
	if err != nil {
		return 0, fmt.Errorf("math undefined")
	}
	solved := math.Tan(num)
	return (solved), nil

}
