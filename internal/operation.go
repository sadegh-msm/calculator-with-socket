package webAppCalculator

import (
	"strconv"
	"strings"
)

// DoMath this function do math calculation and returns the answer
func DoMath(exp string) string {
	if strings.Contains(exp, "+") {
		numbers := strings.Split(exp, "+")

		for _, number := range numbers {
			var res int

			intNum, _ := strconv.Atoi(number)
			res += intNum
			return string(rune(res))
		}
	} else if strings.Contains(exp, "-") {
		numbers := strings.Split(exp, "-")

		for i, number := range numbers {
			var res int

			if i == 1 {
				res, _ = strconv.Atoi(number)
			} else if i > 1 {
				intNum, _ := strconv.Atoi(number)
				res -= intNum
			}
			return string(rune(res))
		}
	} else if strings.Contains(exp, "*") {
		numbers := strings.Split(exp, "*")

		for _, number := range numbers {
			var res = 1

			intNum, _ := strconv.Atoi(number)
			res *= intNum
			return string(rune(res))
		}
	} else if strings.Contains(exp, "/") {
		numbers := strings.Split(exp, "/")

		for i, number := range numbers {
			var res int

			if i == 1 {
				res, _ = strconv.Atoi(number)
			} else if i > 1 {
				intNum, _ := strconv.Atoi(number)
				res /= intNum
			}
			return string(rune(res))
		}
	}
	return "invalid input"
}
