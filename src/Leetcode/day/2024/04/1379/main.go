package main

/*
	傻狗题目，没有 golang 选项
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

/*
class Solution {
public:
    TreeNode* getTargetCopy(TreeNode* original, TreeNode* cloned, TreeNode* target) {
        if (original == nullptr || original == target) {
            return cloned;
        }
        auto left_res = getTargetCopy(original->left, cloned->left, target);
        if (left_res) {
            return left_res; // 已经找到 target，无需递归右子树
        }
        return getTargetCopy(original->right, cloned->right, target);
    }
};
*/
