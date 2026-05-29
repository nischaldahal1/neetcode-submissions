class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        anagram_map = dict()
        for s in strs:
            tuple_val = self.getCount(s)
            if tuple_val not in anagram_map:
                anagram_map[tuple_val] = []
            anagram_map[tuple_val].append(s)
        return list(anagram_map.values())


    def getCount(self, s: str) -> tuple:
        count = [0]*26
        for c in s:
            count[ord(c) - ord('a')] +=1
        return tuple(count)