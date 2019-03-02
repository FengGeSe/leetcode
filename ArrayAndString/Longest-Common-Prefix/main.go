package main

import ()

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	min := len(strs[0])
	for _, s := range strs {
		if len(s) < min {
			min = len(s)
		}
	}

	result := []byte{}
	for i := 0; i < min; i++ {
		c := strs[0][i]
		flag := true
		for _, s := range strs {
			if s[i] != c {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, c)
		} else {
			break
		}

	}
	return string(result)

}
