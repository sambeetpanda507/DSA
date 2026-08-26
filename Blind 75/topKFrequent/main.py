from typing import List
import heapq

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        # Create a hash map
        numsHash: dict[int, int] = {}
        for i in range(len(nums)):
            numsHash[nums[i]] = numsHash.get(nums[i], 0) + 1

        maxHeap: List[int] = []
        for key, val in numsHash.items():
            heapq.heappush(maxHeap, (-val, key))

        result: List[int] = []
        while len(maxHeap) > 0 and k > 0:
            val, num = heapq.heappop(maxHeap)
            result.append(num)
            k = k - 1

        return result

def main():
    # nums = [3, 1,1,1,2,2]
    # k = 2

    # nums = [1]
    # k = 1

    nums = [1,2,1,2,1,2,3,1,3,2]
    k = 2
    print(Solution().topKFrequent(nums, k))


if __name__ == "__main__":
    main()