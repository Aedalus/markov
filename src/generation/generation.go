package generation

import (
	"markov/distribution"
	"math/rand"
	"strings"
	"time"
)

// GetRandomWord retrieves a random next word from a TextDistribution, given a current word
func GetRandomWord(d distribution.TextDistribution, current string) string {
	rand.Seed(time.Now().UnixNano())
	total := 0

	dict := d[current]

	for word := range dict {
		freq := dict[word]
		total += freq
	}

	randNum := rand.Intn(total)

	for word := range dict {
		freq := dict[word]
		total += freq
		if total > randNum {
			return word
		}
	}
	panic("Could not determine word frequency")
}

// GenerateSentence constructs a random sentence from a given TextDistribution
func GenerateSentence(d distribution.TextDistribution) string {
	rand.Seed(time.Now().UnixNano())
	sentence := []string{""}
	for sentence[len(sentence)-1] != "." {
		lastWord := sentence[len(sentence)-1]
		nextWord := GetRandomWord(d, lastWord)
		sentence = append(sentence, nextWord)
	}

	str := strings.Join(sentence, " ")
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, " .", ".")

	return str
}
