/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * class NestedInteger {
 *   public:
 *     // Return true if this NestedInteger holds a single integer, rather than a nested list.
 *     bool isInteger() const;
 *
 *     // Return the single integer that this NestedInteger holds, if it holds a single integer
 *     // The result is undefined if this NestedInteger holds a nested list
 *     int getInteger() const;
 *
 *     // Return the nested list that this NestedInteger holds, if it holds a nested list
 *     // The result is undefined if this NestedInteger holds a single integer
 *     const vector<NestedInteger> &getList() const;
 * };
 */
class NestedIterator {
public:
    NestedIterator(vector<NestedInteger> &nestedList) {
        dfs(nestedList);
    }

    int next() {
        return nums[pos++];
    }

    bool hasNext() {
        return pos < nums.size();
    }
private:
    int pos = 0;
    vector<int> nums;
    
    void dfs(vector<NestedInteger> &nestedList) {
        for (auto &nestedInteger : nestedList) {
            if (nestedInteger.isInteger()) {
                nums.emplace_back(nestedInteger.getInteger());
                continue;
            }
            dfs(nestedInteger.getList());
        }
    }
};

/**
 * Your NestedIterator object will be instantiated and called as such:
 * NestedIterator i(nestedList);
 * while (i.hasNext()) cout << i.next();
 */

