package loopies

func Loopy(s string, n int) string {
	var res string
	for i := 0; i < n; i++ {
		res += s
	}
	return res
}
