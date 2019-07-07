class Item {
public:
    TreeNode *node;
    int level;

    explicit Item(TreeNode *n, int l) : node(n), level(l) {}
};

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    vector<vector<int>> levelOrderBottom(TreeNode *root) {
        vector<vector<int>> ans;
        if (root == nullptr) {
            return ans;
        }
        queue<Item> q;
        q.emplace(root, 0);
        while (!q.empty()) {
            auto item = q.front();
            q.pop();
            while (ans.size() <= item.level) ans.emplace_back();
            ans[item.level].emplace_back(item.node->val);
            if (item.node->left != nullptr) q.emplace(item.node->left, item.level + 1);
            if (item.node->right != nullptr) q.emplace(item.node->right, item.level + 1);
        }

        int l = ans.size();
        for (int i = 0; i < l / 2; ++i) {
            auto tmp = ans[i];
            ans[i] = ans[l - 1 - i];
            ans[l - 1 - i] = tmp;
        }
        return ans;
    }
};

