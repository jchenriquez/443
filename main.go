package main

import (
	"fmt"
	"strconv"
)

func addNum(arr []byte, num int) int {
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

	if len(chars) == 1 {
		return 1
	}
	currCharIndex := 0
	resultLength := 0

	for i := 1; i < len(chars); i++ {

		if chars[i] != chars[currCharIndex] || (i == len(chars)-1 && i-currCharIndex >= 1) {
			var subArr []byte
			var numSteps int

			if i < len(chars)-1 {
				subArr = chars[currCharIndex+1 : i]
				numSteps = i - currCharIndex
			} else {
				subArr = chars[currCharIndex+1 : i+1]
				numSteps = i - currCharIndex
			}

			if numSteps > 1 {
				numLength := addNum(subArr, numSteps)
				currCharIndex += numLength + 1
				chars[currCharIndex] = chars[i]
				resultLength += numLength + 1

			} else {
				currCharIndex += 1
				resultLength++
			}

		}

	}

	return resultLength
}

func main() {
	arr := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'c'}
	length := compress(arr)

	fmt.Printf("compressed %s, length %d", arr, length)
}
