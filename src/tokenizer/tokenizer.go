package tokenizer

import (
	"bufio"
	"strings"
)

// FromString creates an array of tokens from a string
func FromString(text string) []string {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return FromScanner(scanner)
}

// FromScanner creates an array of tokens from a scanner
func FromScanner(scanner *bufio.Scanner) []string {
	scanner.Split(bufio.ScanWords)

	tokens := []string{""}

	for scanner.Scan() {
		next := scanner.Text()

		if strings.HasSuffix(next, ".") {
			tokens = append(tokens, strings.TrimSuffix(next, "."))
			tokens = append(tokens, ".")
			tokens = append(tokens, "")
		} else {
			tokens = append(tokens, next)
		}
	}

	lastToken := tokens[len(tokens)-1]
	if lastToken == "" {
		tokens = tokens[:len(tokens)-1]
	}
	return tokens
}
