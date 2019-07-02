class Solution {
public:
    int removeDuplicates(vector<int> &nums) {
        int len = nums.size();
        if (!len) {
            return 0;
        }
        int pos = 0;
        int prev = nums[0] - 1;
        for (auto &x : nums) {
            if (x == prev) {
                len--;
            } else {
                nums[pos++] = x;
                prev = x;
            }
        }
        return len;
    }
};

