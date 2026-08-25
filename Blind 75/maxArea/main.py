from typing import List

class Solution:
    def maxArea(self, height: List[int]) -> int:
        start = 0
        end = len(height) - 1
        result: int = 0
        while start < end:
            w = end - start
            h = min(height[start], height[end])
            area = w * h
            result = max(area, result)
            if height[start] < height[end]:
                start += 1
            else:
                end -= 1

        return result


def main():
    print("main")
    height = [1,1]
    print(Solution().maxArea(height=height))

if __name__ == "__main__":
    main()