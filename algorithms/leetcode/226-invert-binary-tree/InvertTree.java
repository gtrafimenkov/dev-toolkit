// SPDX-License-Identifier: MIT
// Copyright (c) 2019 Gennady Trafimenkov

public class InvertTree {

    public static void main(String[] args) {
        TreeNode root = new TreeNode(4);
        root.Insert(2);
        root.Insert(7);
        root.Insert(1);
        root.Insert(3);
        root.Insert(6);
        root.Insert(9);
        System.out.println(root);

        Solution sol = new Solution();
        TreeNode inverted = sol.invertTree(root);
        System.out.println(inverted);
    }

}

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) { val = x; }

    public void Insert(int x) {
        if (val < x) {
            right = UpdateLeaf(right, x);
        }
        if (val > x) {
            left = UpdateLeaf(left, x);
        }
    }

    private TreeNode UpdateLeaf(TreeNode leaf, int x) {
        if (leaf == null) {
            return new TreeNode(x);
        }
        leaf.Insert(x);
        return leaf;
    }

    public String toString() {
        return "(" + val
            + " " + left
            + " " + right
            + ")";
    }
}

class Solution {
    public TreeNode invertTree(TreeNode root) {
        if (root == null) {
            return null;
        }
        TreeNode inverted = new TreeNode(root.val);
        inverted.left = invertTree(root.right);
        inverted.right = invertTree(root.left);
        return inverted;
    }
}
