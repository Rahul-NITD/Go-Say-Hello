package choochoo

func Sum(values []int) int {
	res := 0
	for _, value := range values {
		res += value
	}
	return res
}
