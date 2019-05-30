class Solution {
public:
    int thirdMax(vector<int> &nums) {
        long long first = LONG_LONG_MIN;
        long long second = LONG_LONG_MIN;
        long long third = LONG_LONG_MIN;

        for (auto &num : nums) {
            if (num == first || num == second || num == third) {
                continue;
            }
            if (num > first) {
                third = second;
                second = first;
                first = num;
            } else if (num > second) {
                third = second;
                second = num;
            } else if (num > third) {
                third = num;
            }
        }

        return third == LONG_LONG_MIN ? first : third;
    }
};

