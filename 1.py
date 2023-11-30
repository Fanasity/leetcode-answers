class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        for k, v in enumerate(nums):
            b = target - v
            try:
                 i = nums.index(b, k+1)
            except Exception:
                b = 0
            else:
                return [k, i]
        return []