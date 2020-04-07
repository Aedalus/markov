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

	sample.AddTransition("", "foo")
	expected[""]["foo"] = 1

	testEqual(t, sample, expected)

	sample.AddTransition("", "foo")
	expected[""]["foo"] = 2
	testEqual(t, sample, expected)

	sample.AddTransition("", "bar")
	expected[""]["bar"] = 1
	testEqual(t, sample, expected)

	sample.AddTransition("foo", ".")
	expected["foo"] = map[string]int{}
	expected["foo"]["."] = 1
}

func TestFromArray(t *testing.T) {

	sample := FromArray([]string{"", "foo"})
	expected := make(TextDistribution)
	expected.AddTransition("", "foo")

	testEqual(t, sample, expected)

	sample = FromArray([]string{"", "hello", "world", "."})
	expected = make(TextDistribution)
	expected.AddTransition("", "hello")
	expected.AddTransition("hello", "world")
	expected.AddTransition("world", ".")
}
