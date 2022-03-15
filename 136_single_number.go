/*
XOR bitwise operator
a ^ a = 0
a ^ b = b ^ a : commutative
a ^ (b ^ c) = (a ^ b) c) : associative

a ^ a .... ^ a
odd times: a (self)
even times: 0
*/

func singleNumber(nums []int) int {
    answer := 0

    for _, n  := range(nums) {
        answer ^= n
    }
   return answer
}
