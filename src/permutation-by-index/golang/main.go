package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

func main() {
	baseCharset := flag.String("charset", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "Charset")
	length := flag.Int("length", 5, "Length")
	permutation := flag.Int("permutation", 1, "Permutation position (starts with 1)")

	flag.Parse()

	charset := []byte(*baseCharset)
	charsetLength := len(charset)
	if *length > charsetLength {
		fmt.Fprintf(os.Stderr, "The length cannot be more than the count of items in the charset\n")
		return
	}

	maxPermutation := charsetLength - *length + 1
	divisors := make([]int, *length)
	divisors[*length-1] = 1

	for index := 1; index < len(divisors); index++ {
		divisor := charsetLength - *length + index
		targetIndex := *length - index - 1
		divisors[targetIndex] = divisors[targetIndex+1] * divisor
		maxPermutation *= divisor + 1
	}

	if *permutation > maxPermutation {
		fmt.Fprintf(
			os.Stderr,
			"Selected nth permutation exceeds the max based on charset and length (%d)\n",
			maxPermutation,
		)
		return
	}

	codeChars := make([]byte, 0, *length)
	currentPermutation := *permutation
	for _, divisor := range divisors {
		charIndex := int(math.Ceil(float64(currentPermutation)/float64(divisor))) - 1
		currentPermutation = currentPermutation % divisor
		if currentPermutation <= 0 {
			currentPermutation = divisor
		}

		codeChars = append(codeChars, charset[charIndex])
		charset = append(charset[:charIndex], charset[charIndex+1:]...)
	}

	fmt.Printf("Permutation %d: '%s'\n", *permutation, string(codeChars))
}
