package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

func GetPermutationByIndex(charset []byte, length int, index int) (string, error) {
	if index <= 0 {
		return "", fmt.Errorf("index must be greater than 0")
	}

	charsetLength := len(charset)
	if length > charsetLength {
		return "", fmt.Errorf("the length cannot be more than the count of items in the charset (%d)", charsetLength)
	}

	maxPermutation := charsetLength - length + 1
	divisors := make([]int, length)
	divisors[length-1] = 1

	for index := 1; index < len(divisors); index++ {
		divisor := charsetLength - length + index
		targetIndex := length - index - 1
		divisors[targetIndex] = divisors[targetIndex+1] * divisor
		maxPermutation *= divisor + 1
	}

	if index > maxPermutation {
		return "", fmt.Errorf("selected nth permutation exceeds the max based on charset and length (%d)", maxPermutation)
	}

	codeChars := make([]byte, 0, length)
	currentPermutation := index
	for _, divisor := range divisors {
		charIndex := int(math.Ceil(float64(currentPermutation)/float64(divisor))) - 1
		currentPermutation = currentPermutation % divisor
		if currentPermutation <= 0 {
			currentPermutation = divisor
		}

		codeChars = append(codeChars, charset[charIndex])
		charset = append(charset[:charIndex], charset[charIndex+1:]...)
	}

	return string(codeChars), nil
}

func main() {
	baseCharset := flag.String("charset", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "Charset")
	length := flag.Int("length", 5, "Length")
	index := flag.Int("index", 1, "Permutation position (starts with 1)")

	flag.Parse()

	charset := []byte(*baseCharset)
	code, error := GetPermutationByIndex(charset, *length, *index)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	fmt.Printf("Permutation %d: '%s'\n", *index, code)
}
