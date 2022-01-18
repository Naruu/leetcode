class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if len(strs) == 1:
            return strs[0]

        strs.sort(key=lambda s: len(s))
        shortest = strs[0]

        matched = ""
        for i in range(len(shortest) + 1):
            substr = shortest[:i]
            for s in strs[1:]:
                if not s.startswith(substr):
                    return matched
            matched = substr
        return matched
