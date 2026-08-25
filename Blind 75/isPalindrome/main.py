import re

class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = re.sub("[^a-zA-Z0-9]", "", s).lower()
        start = 0
        end = len(s) - 1
        while(start <= end):
            if s[start] != s[end]:
                return False
            else:
                start += 1
                end -= 1

        return True

s = "race a car"
sol = Solution()
print(sol.isPalindrome(s))