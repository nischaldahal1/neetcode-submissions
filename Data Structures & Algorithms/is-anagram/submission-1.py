class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        return self.getCount(s) == self.getCount(t)

    def getCount(self, s) -> tuple:
        hashmap = dict()
        for item in s:
            hashmap[item] = hashmap.get(item, 0) + 1
        return hashmap