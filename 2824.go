func countPairs(nums []int, target int) int {
	count := 0
	for i, num := range nums {
		other := target - num
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < other {
				count++
			}
		}
	}
	return count
}