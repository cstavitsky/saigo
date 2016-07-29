package word_count

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

// Program should read text from an external file
func TestConvertToString(t *testing.T) {
	actual := ConvertToString("example_text.txt")
	assert.Equal(t, actual, "The bot relies on 3D-printed parts, hello")
}

// Program should not include punctuation ,.!?"':; in its counting
func TestStripPunc(t *testing.T) {
	puncString := "The bot? relies! on. 3D-printed parts, hello"
	actual := StripPunc(puncString)
	assert.Equal(t, actual, "The bot relies on 3D-printed parts hello")
}

// Program should provide a correct count of word occurences
// Program should count capital, lowercase as the same word (The/the)
// Whitespace and empty strings should not be counted as words

func TestWordOccurs(t *testing.T) {
	words := []string{"hello","the","when","plastic","hello", "the", "the", "The", "", " "}
	expected := map[string]int{
		"hello": 2,
		"the": 4,
		"when": 1,
		"plastic": 1,
	}
	actual := WordOccurs(words)

	assert.Equal(t, actual, expected)
}