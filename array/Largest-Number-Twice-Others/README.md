
# 题目
在给定的整数数组中，总是存在着最大的一个元素。
判断这个元素是否是其他所有元素的2倍大。
如果两倍大，则返回这个元素的下标，否则返回-1

# 思路
找出数组中最大的两个数，并记录最大数的下标。
如果最大的数比第二大的数的2倍大，则返回下标，否则返回-1.


In a given integer array nums, there is always exactly one largest element.

Find whether the largest element in the array is at least twice as much as every other number in the array.

If it is, return the index of the largest element, otherwise return -1.

Example 1:

Input: nums = [3, 6, 1, 0]
Output: 1
Explanation: 6 is the largest integer, and for every other number in the array x,
6 is more than twice as big as x.  The index of value 6 is 1, so we return 1.
 

Example 2:

Input: nums = [1, 2, 3, 4]
Output: -1
Explanation: 4 isn't at least as big as twice the value of 3, so we return -1.
 

Note:

nums will have a length in the range [1, 50].
Every nums[i] will be an integer in the range [0, 99].
