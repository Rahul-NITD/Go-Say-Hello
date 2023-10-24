package choochoo

func Sum(values []int) int {
	res := 0
	for _, value := range values {
		res += value
	}
	return res
}

func SumAll(numbers ...[]int) []int {
	var ans []int
	for _, nums := range numbers {
		ans = append(ans, Sum(nums))
	}
	return ans
}
