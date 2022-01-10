class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        complement_dict = dict()
        for idx, num in enumerate(nums):
            if complement_dict.get(target-num) is not None:
                return [idx, complement_dict.get(target-num)]
            complement_dict[num] = idx

