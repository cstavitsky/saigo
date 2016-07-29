package main

import (
	"fmt"
	"os"
	"strings"
	"saigo/exercise-001-corpus/count_sort/word_count"
  "sort"
)

type Word struct {
	text string
	numOccurs int
}

type ByOccurrences []Word

func (this ByOccurrences) Len() int {
  return len(this)
}

func (this ByOccurrences) Less(i, j int) bool {
  if this[i].numOccurs == this[j].numOccurs {
    return this[j].text < this[i].text
  } else {
    return this[i].numOccurs < this[j].numOccurs
  }
}

func (this ByOccurrences) Swap(i, j int) {
  this[i], this[j] = this[j], this[i]
}

func main() {
	// Read the file passed in as argument and convert to string
  fileContents := word_count.ConvertToString(os.Args[1:][0])

  // Remove non-word characters from content string
  stripped := word_count.StripPunc(fileContents)

  // Split content string into list of words
  words := strings.Split(stripped, " ")

  // Count and store number of occurrences per word
  wordOccurs := word_count.WordOccurs(words)

  // Sort word occurrences from greatest to least
  var wordObjects []Word
  for k, v := range wordOccurs {
      word := Word{
      text  : k,
      numOccurs: v,
    }
    wordObjects = append(wordObjects, word)
  }

  sort.Sort(sort.Reverse(ByOccurrences(wordObjects)))
  // sort.Reverse(ByOccurrences(wordObjects))

  for _, word := range wordObjects {
    fmt.Println(word.numOccurs, word.text)
  }

}