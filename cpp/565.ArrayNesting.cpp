class Solution {
public:
    int arrayNesting(vector<int> &nums) {
        int ans = 0;
        vector<bool> flags;
        flags.resize(nums.size());
        for (int i = 0; i < nums.size(); ++i) {
            if (flags[i]) continue;
            int num = find(nums, flags, i);
            ans = num > ans ? num : ans;
        }
        return ans;
    }

private:
    static int find(vector<int> &nums, vector<bool> &flags, int start) {
        int n = 0;
        int now = start;
        while (!flags[now]) {
            flags[now] = true;
            n++;
            now = nums[now];
        }
        return n;
    }
};

