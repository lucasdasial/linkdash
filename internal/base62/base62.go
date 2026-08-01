package base62

import "strings"

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

const base = int64(len(alphabet))

// Encode converts a decimal number to base62.
func Encode(n uint64) string {
	if n == 0 {
		return string(alphabet[0])
	}

	var result []byte

	for n > 0 {
		remainder := n % 62
		result = append(result, alphabet[remainder])
		n /= 62
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

// Decode converts a base62 string back to a decimal number.
func Decode(s string) int64 {
	var n int64
	for _, c := range s {
		n = n*base + int64(strings.IndexRune(alphabet, c))
	}
	return n
}
