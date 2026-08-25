from typing import List

class Solution:
    def threeSum(self, nums: list[int]) -> list[list[int]]:
        # Sort the nums array
        nums.sort()
        i = 0 
        result: List[List[int]] = []

        while i < len(nums) - 2:
            j = i + 1
            k = len(nums) - 1
            while j < k:
                sum = nums[i] + nums[j] + nums[k] 
                if sum < 0:
                   j += 1
                elif sum > 0:
                    k -= 1
                else:
                    result.append([nums[i], nums[j], nums[k]])
                    j += 1
                    k -= 1

                    while j < len(nums) and nums[j] == nums[j-1]:
                        j += 1

                    while k >= 0 and nums[k] == nums[k+1]:
                        k -= 1

            i += 1
            while i < len(nums) - 2 and nums[i] == nums[i-1]:
                i += 1

        return result





def main():
    nums = [-1,0,1,2,-1,-4]
    print(Solution().threeSum(nums))



if __name__ == "__main__":
    main()