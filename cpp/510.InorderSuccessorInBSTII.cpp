/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* parent;
};
*/
class Solution {
public:
    Node* inorderSuccessor(Node* node) {
        if (node->right) {
            auto child = node->right;
            while(child->left) {
                child = child->left;
            }
            return child;
        } else {
            auto parent = node->parent;
            while(parent) {
                if (parent->left == node) {
                    return parent;
                } else {
                    node = parent;
                    parent = parent->parent;
                }
            }
        }
        return nullptr;
    }
};

