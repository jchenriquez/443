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
	nextCharPlacement := 0
	resultLength := 0

	for i := 1; i < len(chars); i++ {
		if chars[i] != chars[currCharIndex] || i == len(chars)-1 {
			var distance int
			lastIsDiff := false

			if i == len(chars)-1 {
				fmt.Printf("chars[i] %c, chars[currCharIndex] %c\n", chars[i], chars[currCharIndex])
				if chars[i] != chars[currCharIndex] {
					distance = i - currCharIndex
					resultLength++
					lastIsDiff = true
				} else {
					distance = i + 1 - currCharIndex
				}
			} else {
				distance = i - currCharIndex
			}

			distanceDigitLength := distance

			chars[nextCharPlacement] = chars[currCharIndex]
			if distance > 1 {
				if i-(nextCharPlacement+1) > 0 {
					distanceDigitLength = addNum(chars[nextCharPlacement+1:i], distance)
				} else {
					distanceDigitLength = addNum(chars[nextCharPlacement+1:i+1], distance)
				}
			}

			nextCharPlacement += distanceDigitLength
			currCharIndex = i
			resultLength += distanceDigitLength

			if distance > 1 {
				nextCharPlacement++
				resultLength++
			}

			if i == len(chars)-1 && lastIsDiff {
				chars[nextCharPlacement] = chars[i]
			}

		}

	}

	return resultLength
}

func main() {
	arr := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'a', 'b', 'c'}
	// arr := []byte{'a', 'a', 'a'}
	length := compress(arr)

	fmt.Printf("compressed %s, length %d", arr, length)
}
