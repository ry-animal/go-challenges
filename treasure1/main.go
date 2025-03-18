package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	encodedMap := "20 8 5 0 20 18 5 1 19 21 18 5 0 9 19 0 2 21 18 9 5 4 0 1 20 0 24 25 26 27 28 29"
	decodedMap := decodeTreasureMap(encodedMap)
	fmt.Println("Decoded map:", decodedMap)
}

// Implement the decodeTreasureMap function to decode the secret message
func decodeTreasureMap(encoded string) string {
	tokens := strings.Fields(encoded)

	var decoded strings.Builder // Use a string builder for efficient string concatenation

	for _, token := range tokens {
		num, err := strconv.Atoi(token) // Try to convert token to a number
		if err == nil {                 // If token is a valid number
			if num == 0 {
				decoded.WriteRune(' ') // 0 represents a space
			} else {
				// Handle wrapping: 27->1, 28->2, etc.
				if num > 26 {
					num = ((num - 1) % 26) + 1 // Adjust to wrap around (1-26 range)
				}
				// Convert to corresponding letter (A=1, B=2, etc.)
				decoded.WriteRune(rune('A' + num - 1))
			}
		} else {
			// If not a number, keep the character as is
			decoded.WriteString(token)
		}
	}

	return decoded.String()
}
