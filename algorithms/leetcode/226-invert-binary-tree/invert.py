# SPDX-License-Identifier: MIT
# Copyright (c) 2019 Gennady Trafimenkov

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

    def __str__(self):
        return "({} {} {})".format(self.val, self.left, self.right)

    def insert(self, x):
        def _updateLeaf(leaf, x):
            if leaf == None:
                return TreeNode(x)
            leaf.insert(x)
            return leaf
        if self.val < x:
            self.right = _updateLeaf(self.right, x)
        if self.val > x:
            self.left = _updateLeaf(self.left, x)

class Solution:
    def invertTree(self, root):
        """
        :type root: TreeNode
        :rtype: TreeNode
        """
        if root == None:
            return None
        new_root = TreeNode(root.val)
        new_root.left = self.invertTree(root.right)
        new_root.right = self.invertTree(root.left)
        return new_root

if __name__ == '__main__':
    root = TreeNode(4)
    root.insert(2)
    root.insert(7)
    root.insert(1)
    root.insert(3)
    root.insert(6)
    root.insert(9)
    print(root)
    sol = Solution()
    inverted = sol.invertTree(root)
    print(inverted)
