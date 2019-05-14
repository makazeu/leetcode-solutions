class Solution {
public:
    static bool compare(string &s1, string &s2) {
        if (s1.length() != s2.length()) {
            return s1.length() > s2.length();
        }
        return s1 < s2;
    }

    string findLongestWord(string s, vector<string> &d) {
        string ans;
        sort(d.begin(), d.end(), compare);

        for (auto &word : d) {
            if (word.length() < ans.length()
                || (word.length() == ans.length() && word >= ans)) {
                continue;
            }
            int j = 0;
            for (int i = 0; i < s.length() && j < word.length(); i++) {
                if (s[i] == word[j]) j++;
            }
            if (j == word.length()) {
                ans = word;
            }
        }

        return ans;
    }
};

