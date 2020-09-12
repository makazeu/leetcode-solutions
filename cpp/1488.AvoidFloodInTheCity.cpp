class Solution {
public:
    vector<int> avoidFlood(vector<int>& rains) {
        int n = rains.size();
        vector<int> ans(n);
        set<int> sunny;
        unordered_map <int, int> rainy;

        for (int i = 0; i < n; ++i) {
            int rain = rains[i];
            if (rain == 0) {
                sunny.insert(i);
                continue;
            }
            ans[i] = -1;
            const auto& it = rainy.find(rain);
            if (it == rainy.end()) {
                rainy[rain] = i;
                continue;
            }

            auto res = sunny.upper_bound(it->second);
            if (res == sunny.end()) {
                return vector<int>();
            }
            ans[*res] = rain;
            rainy[rain] = i;
            sunny.erase(res);
        }

        for (auto &day : sunny) {
            ans[day] = 1;
        }
        return ans;
    }
};

