class Item {
public:
    int num;
    int pos;

    explicit Item(int num, int pos): num(num), pos(pos) {}
};

class Solution {
public:
    vector<int> arrayRankTransform(vector<int>& arr) {
        if (arr.empty()) {
            return arr;
        }

        vector<Item> items;
        for (unsigned int i = 0; i < arr.size(); i++) {
            items.emplace_back(arr[i], i);
        }
        sort(items.begin(), items.end(), cmp);

        vector<int> ans(arr.size());
        int rank = 0;
        int lastNum = items[0].num - 1;
        for (auto &item : items) {
            if (item.num != lastNum) {
                rank++;
            }
            lastNum = item.num;
            ans[item.pos] = rank;
        }
        return ans;
    }
private:
    static bool cmp(Item a, Item b) {
        return a.num != b.num ? a.num < b.num : a.pos < b.pos;
    }
};

