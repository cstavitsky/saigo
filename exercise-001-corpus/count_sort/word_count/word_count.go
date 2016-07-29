package word_count

import (
  "io/ioutil"
  "strings"
  "log"
)

func ConvertToString(file string) string {
  // Read file contents as bytes
  bs, err := ioutil.ReadFile(file)
  if err != nil {
    log.Fatal(err)
  }
  // Convert file contents to string
  return string(bs)
}

func StripPunc(str string) string {
  stripped := strings.Replace(str, "(\\W)", "", -1)
  stripped = strings.Replace(stripped, ",", "", -1)
  stripped = strings.Replace(stripped, "?", "", -1)
  stripped = strings.Replace(stripped, "!", "", -1)
  stripped = strings.Replace(stripped, ".", "", -1)
  stripped = strings.Replace(stripped, "\n", " ", -1)
  return stripped
}

func WordOccurs(words []string) map[string]int {
  wordOccurs := make(map[string]int)
  for _, word := range words {
    word = strings.TrimSpace(word)
    if len(word) > 0 {
      word = strings.ToLower(word)
      wordOccurs[word] += 1
    }
  }
  return wordOccurs
}