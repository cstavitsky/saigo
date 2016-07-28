package main

import (
	"fmt"
	"os"
	"strings"
	"saigo/exercise-001-corpus/count_sort/word_count"
)

type Word struct {
	Text string
	Occurs int
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

  fmt.Println(wordOccurs)

  // Sort word occurrences from greatest to least
  // words := []Word
  // for text, occurs := range wordOccurs {
  // 	word := Word{}
  // 	append(words, )
  // }

}