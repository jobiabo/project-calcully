package calcully

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func Power(exp string) (float64, error) {
	fmt.Println("i got here")
	log.Println("Power exp:=", exp)

	cleaned := strings.Split(exp, ")")
	fmt.Println("expecting 2, 5:=", cleaned)
	cleaned1 := strings.Trim(cleaned[0], "x²(")
	fmt.Println(cleaned1)
	// a := cleaned1
	// b := cleaned[1]

	// fmt.Println(cleaned1)
	fmt.Println((cleaned[1]))
	a, err := strconv.ParseFloat(cleaned1, 64)
	if err != nil {
		return 0, err
	}
	fmt.Println(a)

	b, err := strconv.ParseFloat(cleaned[1], 64)
	if err != nil {
		return 0, err

	}
	fmt.Println(b)

	g := (math.Pow(a, b))
	return g, nil
}
