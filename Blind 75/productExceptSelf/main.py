from typing import List

class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        prefix:List[int] = [0] * len(nums)
        suffix: List[int] = [0] * len(nums)
        result: List[int] = [0] * len(nums)
        for i, n in enumerate(nums):
            if i == 0:
                prefix[i] = n
            else:
                prefix[i] = prefix[i-1] * n

        for i in range(len(nums)-1, -1, -1):
            if i == len(nums) - 1:
                suffix[i] = nums[i]
            else:
                suffix[i] = suffix[i+1] * nums[i]


        for i in range(len(nums)):
            if i == 0:
                result[i] = suffix[i+1]
            elif i == len(nums) - 1:
                result[i] = prefix[i-1]
            else:
                result[i] = prefix[i-1] * suffix[i+1]

        return result

def main():
    print("main")
    nums = [-1,1,0,-3,3]
    sol = Solution()
    result = sol.productExceptSelf(nums)
    print(f"result: {result}")


if __name__ == "__main__":
    main()