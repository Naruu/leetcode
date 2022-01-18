PAIR = {
    "(": -1,
    ")": 1,
    "{": -2,
    "}": 2,
    "[": -3,
    "]": 3
}


class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        for string in s:
            if PAIR[string] < 0:
                stack.append(string)
                continue
            if len(stack) < 1:
                return False
            latest = stack.pop()
            if PAIR[latest] + PAIR[string] != 0:
                return False

        if len(stack) > 0:
            return False

        return True
