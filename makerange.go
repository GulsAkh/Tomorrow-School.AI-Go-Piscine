package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int{}
	}

	answer := make([]int, max-min)
	for i := 0; i < len(answer); i++ {
		answer[i] = min + i
	}
	return answer
}
