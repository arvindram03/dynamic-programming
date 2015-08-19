package dynamicprogramming

func FindMaxSeq(seq []int) int {
	n := len(seq)
	if n == 1 {
		return 1
	}

	prev := 0
	ls := make([]int, n)
	ls[prev] = 1

	for curr := 1; curr < n; curr++ {
		ls[curr] = 1
		if seq[curr] > seq[prev] {
			ls[curr] += ls[prev]
		}
		prev = curr
	}

	return max(ls)
}

func max(s []int) (max int) {
	for i := 0; i < len(s); i++ {
		if s[i] >= max {
			max = s[i]
		}
	}
	return
}
