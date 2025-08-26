package iterations

import "strings"

const RepeationValue = 5

func Repeat(value string) string {
	var answer strings.Builder
	for i := 0; i < RepeationValue; i++ {
		answer.WriteString(value)
	}
	return answer.String()
}
