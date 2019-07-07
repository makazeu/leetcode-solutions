class Solution {
public:
    bool detectCapitalUse(string word) {
        return
                all_of(word.begin(), word.end(), [](char c) { return islower(c); })
                || all_of(word.begin(), word.end(), [](char c) { return isupper(c); })
                || (word.length() > 0 && isupper(word[0]) &&
                    all_of(word.begin() + 1, word.end(), [](char c) { return islower(c); }));

    }
};

