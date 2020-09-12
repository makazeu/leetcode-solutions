class Solution {
public:
    vector<int> smallerNumbersThanCurrent(vector<int>& nums) {
        multiset<int> s(nums.begin(), nums.end());
        vector<int> ans;
        for (auto &num : nums) {
            ans.emplace_back(distance(s.begin(), s.lower_bound(num)));
        }
        return ans;
    }
};

