package main

import ()

func AddBinary(a, b string) string {

	n := len(b)
	if len(a) > len(b) {
		n = len(a)
	}
	flag := false
	result := make([]byte, n+1)
	for i := 0; i <= n; i++ {

		var x, y byte
		ia := len(a) - 1 - i
		if ia >= 0 {
			x = a[ia]
		}

		ib := len(b) - 1 - i
		if ib >= 0 {
			y = b[ib]
		}

		sum := 0
		if flag {
			sum++
		}

		if x == byte('1') {
			sum++
		}

		if y == byte('1') {
			sum++
		}

		switch sum {
		case 0:
			result[len(result)-1-i] = byte('0')
			flag = false
		case 1:
			result[len(result)-1-i] = byte('1')
			flag = false
		case 2:
			result[len(result)-1-i] = byte('0')
			flag = true
		case 3:
			result[len(result)-1-i] = byte('1')
			flag = true

		}

	}

	if result[0] == byte('0') {
		result = result[1:]
	}

	return string(result)

}
