from functools import cmp_to_key

class Solution:
    def sortedSquares(self, A: List[int]) -> List[int]:
        A.sort(key=cmp_to_key(lambda x, y: abs(x) - abs(y)))
        return [x * x for x in A]
