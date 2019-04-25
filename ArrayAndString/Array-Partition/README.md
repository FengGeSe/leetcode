
# 题目
Given an array of **2n** integers, your task is to group these integers into **n** pairs of integer, say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.



**Example 1:**

Input: 
```
[1,4,3,2]
```
Output: 
> 4

Explanation:

​	n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3, 4).



**Note:**

1. **n** is a positive integer, which is in the range of [1, 10000].
2. All the integers in the array will be in the range of [-10000, 10000].



```
func arrayPairSum(nums []int) int {
    
		var backet [20001]int
	for _, val := range nums {
		backet[10000+val]++
	}
	count := 0
	sum := 0
	for index := 0; index < 20001; index++ {
		for backet[index] > 0 {
			count++
			if count&1 == 1 {
				sum += index - 10000
			}
			backet[index]--
		}
	}
    return sum
}
```

