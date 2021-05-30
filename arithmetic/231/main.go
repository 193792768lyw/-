package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(0))
}

func isPowerOfTwo(n int) bool {

	if n < 1 {
		return false
	}
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			return false
		}

	}
	return true

}
