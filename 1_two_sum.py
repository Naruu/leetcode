class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        idx_dict = dict()
        for idx, num in enumerate(nums):
            pair = target-num
            if idx_dict.get(pair) is not None:
                return [idx, idx_dict.get(pair)]
            idx_dict[num] = idx
