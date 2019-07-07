/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};
*/
class Solution {
public:
    vector<int> preorder(Node *root) {
        vector<int> ans;
        dfs(root, ans);
        return ans;
    }

    void dfs(Node *node, vector<int> &result) {
        if (node == nullptr) return;
        result.emplace_back(node->val);
        for (auto &child : node->children) {
            dfs(child, result);
        }
    }
};

