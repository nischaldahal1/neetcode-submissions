class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        map = dict()
        for i, item in enumerate(nums):
            if (target - item) in map:
                return [map[target - item],i]
            else:
                map[item] = i
        return [-1,-1]