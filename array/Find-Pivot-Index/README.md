# 算法
找出一个数组中的平衡点的下标。
平衡点左边的数之和等于平衡点右边数之和。
# 思路
假定pivot点初始是0
那么左边数之和leftSum = 0
右边数之和rightSum = sum(nums)
pivot点坐标右移的时候rightSum减去这个点的值, leftSum加上这个点的值。

# Introduction 
Given an array of integers nums, write a method that returns the "pivot" index of this array.

We define the pivot index as the index where the sum of the numbers to the left of the index is equal to the sum of the numbers to the right of the index.

If no such index exists, we should return -1. If there are multiple pivot indexes, you should return the left-most pivot index.

Example 1:
```
Input: 
nums = [1, 7, 3, 6, 5, 6]
Output: 3
Explanation: 
The sum of the numbers to the left of index 3 (nums[3] = 6) is equal to the sum of numbers to the right of index 3.
Also, 3 is the first index where this occurs.
```
Example 2:
```
Input: 
nums = [1, 2, 3]
Output: -1
Explanation: 
There is no index that satisfies the conditions in the problem statement.
```
Note:

The length of nums will be in the range [0, 10000].
Each element nums[i] will be an integer in the range [-1000, 1000].

