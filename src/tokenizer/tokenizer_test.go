package tokenizer

import (
	"reflect"
	"testing"
)

func testTokenString(t *testing.T, tokens, expected []string) {
	if reflect.DeepEqual(tokens, expected) == false {
		t.Log("Tokens other than expected")
		t.Logf("%+v", tokens)
		t.Logf("%+v", expected)
		t.Fail()
	}
}
func TestStartNewSentenceWithEmptyString(t *testing.T) {
	testTokenString(t,
		FromString("Hello World"),
		[]string{
			"",
			"Hello",
			"World",
		},
	)
}

func TestSplitPunctuation(t *testing.T) {
	testTokenString(t,
		FromString("Hello."),
		[]string{
			"",
			"Hello",
			".",
		},
	)

	testTokenString(t,
		FromString("Hello. World."),
		[]string{
			"",
			"Hello",
			".",
			"",
			"World",
			".",
		},
	)
}
