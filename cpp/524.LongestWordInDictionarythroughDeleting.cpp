class Solution {
public:
    string findLongestWord(string s, vector<string> &d) {
        string ans;

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

