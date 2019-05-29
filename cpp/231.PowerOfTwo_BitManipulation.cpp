class Solution {
public:
    bool isPowerOfTwo(int n) {
        return n && !(n & (long long) n - 1);
    }
};

