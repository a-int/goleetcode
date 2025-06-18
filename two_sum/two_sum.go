package twosum

func twoSum(nums []int, target int) [2]int {
	var pairs map[int]int = make(map[int]int)
	for idx, val := range nums {
		_, ok := pairs[val]

		if ok {
			return [2]int{pairs[val], idx}
		} else {
			pairs[target-val] = idx
		}
	}

	return [2]int{0, 0}
}
