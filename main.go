package main

import (
	"fmt"
	"strconv"
)

// Generate sequential CHAR from NUM
// Eg. 001 -> AAB and so on
func main() {
	fmt.Println(generateCharFromNum(21, 6))
}

func generateCharFromNum(start, digit int) string {
	var res string
	numForm := "%0" + strconv.Itoa(digit) + "d"
	n := fmt.Sprintf(numForm, start+1)
	for _, v := range n {
		res += toChar(v)
	}
	return res
}

func toChar(i rune) string {
	return string('A' - 0 + (int(i) - '0'))
}
