class Solution {
public:
    bool canThreePartsEqualSum(vector<int> &A) {
        int sum = accumulate(A.begin(), A.end(), 0);
        if (sum % 3 != 0) {
            return false;
        }
        sum /= 3;

        int now = 0;
        int have = 0;
        for (auto &x : A) {
            now += x;
            if (now == sum) {
                have++;
                now = 0;
                if (have == 3) {
                    return true;
                }
            }
        }
        return false;
    }
};

