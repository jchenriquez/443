package main

import "strconv"

func addNum (arr[] byte, num int) int {
	numStr := strconv.Itoa(num)

	for i := 0; i < len(numStr); i++ {
		arr[i] = numStr[i]
	}

	return len(numStr)
}

func compress(chars []byte) int {

	if len(chars) == 0 {
		return 0
	}

	nextWriteIndex := 1
	currCharIndex := 0

	for i := 1; i < len(chars); i++ {

		if chars[i] != chars[currCharIndex] {

		}
	}
}

func main() {

}
