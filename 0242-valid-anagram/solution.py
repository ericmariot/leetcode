class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        f_word, s_word = {}, {}

        for i in range(len(s)):
            f_word[s[i]] = 1 + f_word.get(s[i], 0)
            s_word[t[i]] = 1 + s_word.get(t[i], 0)
        
        return f_word == s_word
