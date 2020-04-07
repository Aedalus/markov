package distribution

import "math/rand"

// GetRandWord returns a random next word given the current word.
// This value is determined from the transition weights.
func (d TextDistribution) GetRandWord(current string) string {

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
