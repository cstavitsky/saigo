package word_sort

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

// Program should build word objects with word text and number of occurrences
func TestBuildWords(t *testing.T) {
	wordData := map[string]int{
		"hello": 2,
		"the": 4,
		"when": 1,
		"plastic": 1,
	}

	expected := []Word{
		{Text: "hello", NumOccurs: 2},
		{Text: "the", NumOccurs: 4},
		{Text: "when", NumOccurs: 1},
		{Text: "plastic", NumOccurs: 1},
	}

	assert.Equal(t, BuildWords(wordData), expected)
}

//Program should sort words by number of occurrences
//If there is a tie in occurrences, program should sort tied words alphabetically
func TestSortWords(t *testing.T) {
	wordObjects := []Word{
		{Text: "hello", NumOccurs: 2},
		{Text: "the", NumOccurs: 4},
		{Text: "when", NumOccurs: 1},
		{Text: "plastic", NumOccurs: 1},
	}

	expected := []Word{
		{Text: "the", NumOccurs: 4},
		{Text: "hello", NumOccurs: 2},
		{Text: "plastic", NumOccurs: 1},
		{Text: "when", NumOccurs: 1},
	}

	SortWords(wordObjects)

	assert.Equal(t, expected, wordObjects)
}
