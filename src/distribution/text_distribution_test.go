package distribution

import (
	"reflect"
	"testing"
)

func testEqual(t *testing.T, sample, expected TextDistribution) {

	if reflect.DeepEqual(sample, expected) {

	} else {
		t.Log("Error: The following should match")
		t.Logf("%+v", sample)
		t.Logf("%+v", expected)
		t.Fail()
	}
}

func TestAddTransition(t *testing.T) {
	sample := make(TextDistribution)
	expected := make(TextDistribution)
	expected[""] = map[string]int{}

	sample.AddTransition("", "foo", 1)
	expected[""]["foo"] = 1

	testEqual(t, sample, expected)

	sample.AddTransition("", "foo", 1)
	expected[""]["foo"] = 2
	testEqual(t, sample, expected)

	sample.AddTransition("", "bar", 1)
	expected[""]["bar"] = 1
	testEqual(t, sample, expected)

	sample.AddTransition("foo", ".", 1)
	expected["foo"] = map[string]int{}
	expected["foo"]["."] = 1
	testEqual(t, sample, expected)

	sample.AddTransition("Hello", "World", 10)
	expected["Hello"] = map[string]int{}
	expected["Hello"]["World"] = 10
	testEqual(t, sample, expected)
}

func TestFromArray(t *testing.T) {

	sample := FromArray([]string{"", "foo"})
	expected := make(TextDistribution)
	expected.AddTransition("", "foo", 1)

	testEqual(t, sample, expected)

	sample = FromArray([]string{"", "hello", "world", "."})
	expected = make(TextDistribution)
	expected.AddTransition("", "hello", 1)
	expected.AddTransition("hello", "world", 1)
	expected.AddTransition("world", ".", 1)
}

func TestCombine(t *testing.T) {

	sampleA := FromArray([]string{})
	sampleB := FromArray([]string{})
	expected := make(TextDistribution)
	testEqual(t, sampleA, expected)
	testEqual(t, sampleB, expected)

	sampleA.AddTransition("A", "B", 5)
	sampleA.AddTransition("A", "C", 1)
	sampleA.AddTransition("A", "D", 1)

	sampleB.AddTransition("A", "B", 1)
	sampleB.AddTransition("A", "C", 2)
	sampleB.AddTransition("A", "E", 2)

	sampleC := sampleA.Combine(sampleB)

	expected = make(TextDistribution)
	expected.AddTransition("A", "B", 6)
	expected.AddTransition("A", "C", 3)
	expected.AddTransition("A", "D", 1)
	expected.AddTransition("A", "E", 2)
	testEqual(t, sampleC, expected)
}
