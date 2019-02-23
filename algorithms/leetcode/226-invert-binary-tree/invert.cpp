// SPDX-License-Identifier: MIT
// Copyright (c) 2019 Gennady Trafimenkov

#include <iostream>
#include <string>
#include <sstream>

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}

    std::string GetString() {
        std::ostringstream ss;
        ss << '(' << this->val
            << ' ' << (this->left ? this->left->GetString() : "null")
            << ' ' << (this->right ? this->right->GetString() : "null")
            << ')';
        return ss.str();
    }

    void Insert(int x) {
        if (this->val < x) {
            this->right = UpdateLeaf(this->right, x);
        }
        if (this->val > x) {
            this->left = UpdateLeaf(this->left, x);
        }
    }

    ~TreeNode() {
        if (left) delete left;
        if (right) delete right;
    }

    private:

    static TreeNode* UpdateLeaf(TreeNode* leaf, int x) {
        if (!leaf) {
            return new TreeNode(x);
        }
        leaf->Insert(x);
        return leaf;
    }
};

class Solution {
public:
    TreeNode* invertTree(TreeNode* root) {
        if (root) {
            TreeNode* inverted = new TreeNode(root->val);
            inverted->left = invertTree(root->right);
            inverted->right = invertTree(root->left);
            return inverted;
        }
        return NULL;
    }
};

int main() {
    TreeNode* root = new TreeNode(4);
    root->Insert(2);
    root->Insert(7);
    root->Insert(1);
    root->Insert(3);
    root->Insert(6);
    root->Insert(9);
    std::cout << root->GetString() << std::endl;

    Solution sol;
    TreeNode* inverted = sol.invertTree(root);
    std::cout << inverted->GetString() << std::endl;

    delete inverted;
    delete root;

    return 0;
}
