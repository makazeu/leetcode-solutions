class Solution {
public:
    int removeElement(vector<int> &nums, int val) {
        int pos = 0;
        int len = nums.size();
        for (auto &x : nums) {
            if (x == val) {
                len--;
            } else {
                nums[pos++] = x;
            }
        }
        return len;
    }
};

