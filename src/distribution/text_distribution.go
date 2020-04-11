package distribution

// TextDistribution represents a probability of next tokens
type TextDistribution map[string]map[string]int

// FromArray takes an array of strings, and returns a new TextDistribution.
func FromArray(tokens []string) TextDistribution {
	distribution := make(TextDistribution)
	for i, current := range tokens {
		if i == len(tokens)-1 {
			break
		}
		next := tokens[i+1]
		distribution.AddTransition(current, next, 1)
	}
	return distribution
}

// AddTransition increments the weight of a new or existing markov transition.
// Assuming the given state is "Hello" and the transition is "World",
// conceptually this would be represented as ["Hello"]["World"] == 1
// Calling this multiple times with the same values will increment the weight of the transition.
func (d TextDistribution) AddTransition(state, transition string, count int) {
	if d[state] == nil {
		d[state] = make(map[string]int)
	}

	distState := d[state]
	distState[transition] += count
}

// Combine creates a new TextDistribution that is the combined result of
// the original TextDistribution with another TextDistribution
func (d TextDistribution) Combine(other TextDistribution) TextDistribution {
	newDist := make(TextDistribution)

	for nodeA := range d {
		for edgeA := range d[nodeA] {
			countA := d[nodeA][edgeA]
			newDist.AddTransition(nodeA, edgeA, countA)
		}
	}

	for nodeB := range other {
		for edgeB := range other[nodeB] {
			countB := other[nodeB][edgeB]
			newDist.AddTransition(nodeB, edgeB, countB)
		}
	}

	return newDist
}
