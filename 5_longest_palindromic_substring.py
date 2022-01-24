class Solution:
    def expandPalindrome(self, s, left, right):
        while left > 0 and right < len(s) - 1:
            if s[left-1] != s[right + 1]:
                return left, right + 1
            left -= 1
            right += 1
        return left, right+1

    def longestPalindrome(self, s: str) -> str:
        answer = ""

        for i in range(len(s)):
            start, end = self.expandPalindrome(s, i, i)
            if len(answer) < (end - start):
                answer = s[start:end]

        for i in range(len(s)-1):
            if s[i] == s[i+1]:
                start, end = self.expandPalindrome(s, i, i+1)
                if len(answer) < (end - start):
                    answer = s[start:end]

        return answer
