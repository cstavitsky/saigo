package word_sort

type Word struct {
	Text string
	NumOccurs int
}

type ByOccurrences []Word

func (this ByOccurrences) Len() int {
  return len(this)
}

func (this ByOccurrences) Less(i, j int) bool {
  if this[i].NumOccurs == this[j].NumOccurs {
    return this[j].Text < this[i].Text
  } else {
    return this[i].NumOccurs < this[j].NumOccurs
  }
}

func (this ByOccurrences) Swap(i, j int) {
  this[i], this[j] = this[j], this[i]
}