#
# @lc app=leetcode id=100 lang=python3
#
# [100] Same Tree
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from typing import Optional


class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        def helper(node1, node2) -> bool:
            if (not node1 and node2) or (node1 and not node2):
                return False
            if not node1 and not node2:
                return True
            if node1.val != node2.val:
                return False
            return helper(node1.left, node2.left) and helper(node1.right, node2.right)
        return helper(p, q)

# @lc code=end
