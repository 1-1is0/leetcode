class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        char_list = []
        count = 0
        for a in s:
            try:
                a_index = char_list.index(a)
                char_list = char_list[a_index+1:]
                char_list.append(a)
            except ValueError:
                a_index = None
                char_list.append(a)
            
            if count < len(char_list):
                count = len(char_list)
        return count

# s = "abcabcbb"
# s = "bbbbb"
# s = "pwwkew"
# s = " "
# s = "dvdf"
sol = Solution()
print(sol.lengthOfLongestSubstring(s))