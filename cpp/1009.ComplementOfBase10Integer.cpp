class Solution {
public:
    int bitwiseComplement(int N) {
        if (N == 0) return 1;
        int d = (int) log2(N) + 1;
        return N ^ ((1 << d) - 1);
    }
};

