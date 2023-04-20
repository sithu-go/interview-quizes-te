package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(getNumber("Var-----___1_int"))
	fmt.Println(getNumber("Q2q-q"))
	fmt.Println(getNumber("eef3243**s"))

}

func getNumber(input string) string {
	for _, v := range input {
		if unicode.IsNumber(v) {
			return string(v)
		}
	}
	return ""
}
