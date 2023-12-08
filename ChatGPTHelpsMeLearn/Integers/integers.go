package integers

func SumUpto(n int) int {
	if n <= 0 {
		return 0
	}
	s := n * (n + 1) / 2
	return s
}
