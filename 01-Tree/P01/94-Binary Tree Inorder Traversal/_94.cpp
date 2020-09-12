//
// Created by Bridge_Wang on 2020/9/12.
//
#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <unordered_map>
using namespace std;

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
 struct TreeNode {
     int val;
     TreeNode *left;
     TreeNode *right;
     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 };
class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> res;
        dfs(root, &res);
        return res;
    }

    void dfs(TreeNode *root, vector<int>* res) {
        if (root != NULL) {
            dfs(root->left, res);
            res->push_back(root->val);
            dfs(root->right, res);
        }
    }
};