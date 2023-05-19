package main

import "fmt"

func demo() {
	var n string
	fmt.Scanln(&n)
	var ans = "0."

	for i := 2; i < len(n)-1; i++ {
		var index = 0
		temp := n[i]
		for j := i + 1; j < len(n); j++ {
			if n[j] > temp {
				temp = n[j]
				index = j
			}
		}
		// 字符串拼接 ,temp转
		ans = ans + string(rune(temp))
		// 调整i
		i += index - i
	}
	fmt.Println(ans)
}
