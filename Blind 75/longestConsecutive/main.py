from typing import List

class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        # Store the nums in set
        numSet: set[int] = set(nums)
        result: int = 0
        curr: int = 0
        for num in numSet:
            start = num
            # Check start of sequence
            if start - 1 not in numSet:
                while start in numSet:
                    curr += 1
                    start += 1

            result = max(result, curr)
            curr = 0

        return result

def main():
    nums = [1,0,1,2]
    sol = Solution()
    result = sol.longestConsecutive(nums)
    print(f"result: {result}")

if __name__ == "__main__":
    main()