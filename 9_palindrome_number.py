class Solution1:
    def isPalindrome(self, x: int) -> bool:
        str_x = str(x)
        len_string = len(str_x)
        for i in range(len_string//2):
            if str_x[i] != str_x[len_string - i - 1]:
                return False
        return True

class Solution2:
    def isPalindrome(self, x: int) -> bool:
        if str(x) != str(x)[::-1]:
            return False
        return True