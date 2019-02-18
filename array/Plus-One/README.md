# 题目
用一个非空的数组代表一个非负整数，然后+1。
非负整数的最高位在数组的头部，每个数组元素代表一位数。
必须保证数组不能以0开头。除了0本身。


# 思路
尾数+1，当+1后要进位，则继续往前加。当前为变为0






Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
