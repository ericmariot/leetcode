class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        nums_map = {}

        for i in range(len(nums)):
            t = target - nums[i]
            if t in nums_map:
                return [nums_map[t], i]
            nums_map[nums[i]] = i

        return []


