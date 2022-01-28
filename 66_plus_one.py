class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        answer = []
        carry = 1
        for digit in digits[::-1]:
            s = (carry + digit)
            answer.append(s % 10)
            carry = s // 10
        if carry:
            answer.append(carry)
        return answer[::-1]
