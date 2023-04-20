package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(isValidIP("172.255.254.1")) // true

	fmt.Println(isValidIP("172.316.254.1")) // false
	fmt.Println(isValidIP("0.1.1.256"))     // false
	fmt.Println(isValidIP("1.1.1.1a"))      // false
	fmt.Println(isValidIP("1"))             // false
	fmt.Println(isValidIP("64.233.16.00"))  // false
	fmt.Println(isValidIP("7283728"))       // false

}

func isValidIP(ip string) bool {
	if strings.Count(ip, ".") == 3 {
		nos := strings.Split(ip, ".")
		for _, v := range nos {
			no, err := strconv.Atoi(v)
			if err != nil {
				return false
			}
			if no < 0 || no > 255 {
				return false
			}

			if no == 0 {
				if strings.Count(v, "0") > 1 {
					return false
				}
			}

		}
		return true
	}

	return false
}
