from typing import List
# quesion link: https://leetcode.com/problems/two-sum/description/
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in range(0, len(nums)):
            n1 = nums[i]
            for j in range(i+1, len(nums)):
                n2 = nums[j]
                if n1 + n2 == target:
                    return [i, j]
        return [0, 0]




nums = [2,7,11,15]
target = 9
sol = Solution()
print(sol.twoSum(nums, target))
