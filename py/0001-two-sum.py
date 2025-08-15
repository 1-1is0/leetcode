from typing import List
# question link: https://leetcode.com/problems/two-sum/description/
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in range(0, len(nums)):
            n1 = nums[i]
            for j in range(i+1, len(nums)):
                n2 = nums[j]
                if n1 + n2 == target:
                    return [i, j]
        return [0, 0]

    def twoSumHashMap(self, nums: List[int], target: int) -> List[int]:
        hashmap = {}
        for i in range(len(nums)):
            hashmap[nums[i]] = i
        # print(hashmap)
        for i in range(0, len(nums)):
            other_half_val = target - nums[i]
            if other_half_val in hashmap and hashmap[other_half_val] != i:
                return [i, hashmap[other_half_val]]
        return [0, 0]




nums = [2,7,11,15]
target = 9

# nums = [3, 2, 4]
# target = 6
sol = Solution()
# print(sol.twoSum(nums, target))
print(sol.twoSumHashMap(nums, target))
