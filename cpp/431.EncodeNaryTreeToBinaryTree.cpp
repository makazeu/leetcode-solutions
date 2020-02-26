/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val) {
        val = _val;
    }

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};
*/
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Codec {
public:

    // Encodes an n-ary tree to a binary tree.
    TreeNode* encode(Node* root) {
        if (!root) {
            return nullptr;
        }
        auto newNode = new TreeNode(root->val);
        if (root->children.empty()) {
            return newNode;
        }
        newNode->left = encode(root->children[0]);

        auto currentNode = newNode->left;
        for (int i = 1; i < root->children.size(); i++) {
            currentNode->right = encode(root->children[i]);
            currentNode = currentNode->right;
        }
        return newNode;
    }

    // Decodes your binary tree to an n-ary tree.
    Node* decode(TreeNode* root) {
        if (!root) {
            return nullptr;
        }
        auto newNode = new Node(root->val);
        if (!root->left) {
            return newNode;
        }
        newNode->children.emplace_back(decode(root->left));

        auto currentNode = root->left;
        while (currentNode->right) {
            newNode->children.emplace_back(decode(currentNode->right));
            currentNode = currentNode->right;
        }
        return newNode;
    }
};

// Your Codec object will be instantiated and called as such:
// Codec codec;
// codec.decode(codec.encode(root));
