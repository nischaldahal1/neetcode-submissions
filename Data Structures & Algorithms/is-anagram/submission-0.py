class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        return self.getKey(s) == self.getKey(t)

    def getKey(self, s) -> tuple:
        key_list = [0] * 26
        for c in s:
           key_list[ord(c)- ord('a')] +=1
        return tuple(key_list)