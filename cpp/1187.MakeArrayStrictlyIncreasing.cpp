class Solution {
public:
    int makeArrayIncreasing(vector<int> &arr1, vector<int> &arr2) {
        sort(arr2.begin(), arr2.end());
        auto it = unique(arr2.begin(), arr2.end());
        arr2.erase(it, arr2.end());

        int dp[arr1.size()][arr2.size() + 1];
        dp[0][0] = 0;
        for (int i = 1; i <= arr2.size(); ++i) {
            dp[0][i] = 1;
        }
        for (int i = 1; i < arr1.size(); ++i) {
            for (int j = 0; j <= arr2.size(); ++j) {
                dp[i][j] = INT32_MAX - 1;
            }
        }

        for (int i = 1; i < arr1.size(); ++i) {
            // no operation
            if (arr1[i - 1] < arr1[i]) {
                dp[i][0] = dp[i - 1][0];
            }
            for (int j = 0; j < arr2.size() && arr2[j] < arr1[i]; ++j) {
                dp[i][0] = min(dp[i][0], dp[i - 1][j + 1]);
            }

            // do operation
            for (int j = 0; j < arr2.size(); ++j) {
                if (j > 0) {
                    dp[i][j + 1] = min(dp[i][j + 1], dp[i - 1][j] + 1);
                }
                if (arr1[i - 1] < arr2[j]) {
                    dp[i][j + 1] = min(dp[i][j + 1], dp[i - 1][0] + 1);
                }
            }
        }

        int ans = INT32_MAX - 1;
        for (int i = 0; i <= arr2.size(); ++i) {
            ans = min(ans, dp[arr1.size() - 1][i]);
        }

        return ans == INT32_MAX - 1 ? -1 : ans;
    }
};

