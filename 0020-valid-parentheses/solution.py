class Solution:
    def isValid(self, s: str) -> bool:
        opcl = dict(('()', '[]', '{}'))
        stack = []

        for c in s:
            if c in "([{":
                stack.append(c)
            elif len(stack) == 0 or c != opcl[stack.pop()]:
                return False


        return len(stack) == 0
