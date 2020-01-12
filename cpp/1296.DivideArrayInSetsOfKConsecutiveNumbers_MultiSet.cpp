class Solution {
public:
    bool isPossibleDivide(vector<int> &nums, int k) {
        int len = nums.size();
        if (len % k != 0) {
            return false;
        }

        multiset<int> s;
        for (auto &e : nums) {
            s.insert(e);
        }

        while (!s.empty()) {
            auto p = s.begin();
            int num = *p;
            s.erase(p);
            for (int i = 1; i <= k - 1; ++i) {
                p = s.upper_bound(num);
                if (*p != num + 1) {
                    return false;
                }
                s.erase(p);
                num++;
            }
        }
        return true;
    }
};

