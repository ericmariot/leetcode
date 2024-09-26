class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = [c.lower() for c in s if c.isalnum()]
        left, right = 0, len(s)-1

        while left < right:
            if s[left] != s[right]:
                return False
            left = left + 1
            right = right - 1
        
        
        return True

