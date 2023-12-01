package aoc

import (
	"fmt"
	"regexp"
)

func printStruct(s any) {
	fmt.Printf("%+v\n", s)
}

func getNumbersFromString(s string) []string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAllString(s, -1)
}
