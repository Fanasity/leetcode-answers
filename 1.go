func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}
	for i, x := range nums {
		y := target - x
		for j, m := range nums {
			if i != j && m == y {
				return []int{i, j}
			}
		}
	}
	return []int{}
}