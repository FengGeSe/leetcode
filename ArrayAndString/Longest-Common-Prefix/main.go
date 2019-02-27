package main

import ()

func StrStr(haystack, needle string) int {

	result := -1
	if needle == "" {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	start := 0
	for i := 0; i < len(haystack); i++ {
		start = i
		if haystack[i] == needle[0] && len(haystack)-i >= len(needle) {
			flag := true
			t := start
			for j, _ := range needle {
				if haystack[t] != needle[j] {
					flag = false
					break
				}
				t++
			}
			if flag {
				result = start
				break
			}
		}

	}
	return result

}
