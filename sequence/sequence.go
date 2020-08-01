package sequence

type Sequence []int

func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	if s[i] < s[j] {
		return true
	} else {
		return false
	}
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var Seq = Sequence{111, 2, 11, 939, 31, 2312, 124, 92, 515, 67, 43, 818, 203, 792, 399, 2, 11, 24, 55, 211, 79, 12, 84, 34, 29, 90, 103,
	110, 234, 765, 124, 555, 135}
