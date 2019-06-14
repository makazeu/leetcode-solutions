class Solution {
public:
    string customSortString(string S, string T) {
        int order[26];
        for (int i = 0; i < S.size(); i++) {
            order[S[i] - 'a'] = i;
        }

        auto cmp_func = [=](char a, char b) {
            return order[a - 'a'] < order[b - 'a'];
        };
        sort(T.begin(), T.end(), cmp_func);
        return T;
    }
};

