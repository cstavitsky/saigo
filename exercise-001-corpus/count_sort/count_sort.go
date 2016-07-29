package main

import (
	"fmt"
	"os"
	"strings"
	"saigo/exercise-001-corpus/count_sort/word_count"
  "saigo/exercise-001-corpus/count_sort/word_sort"
)

func main() {
	// Read the file passed in as argument and convert to string
  fileContents := word_count.ConvertToString(os.Args[1:][0])

  // Remove non-word characters from content string
  stripped := word_count.StripPunc(fileContents)

  // Split content string into list of words
  words := strings.Split(stripped, " ")

  // Count and store number of occurrences per word
  wordOccurs := word_count.WordOccurs(words)

  // Create Word objects
  wordObjects := word_sort.BuildWords(wordOccurs)

  word_sort.SortWords(wordObjects)

  // Print words and occurrences to console
  for _, word := range wordObjects {
    fmt.Println(word.NumOccurs, word.Text)
  }

}