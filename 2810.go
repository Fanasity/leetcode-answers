package main

import "fmt"

func finalString(s string) string {
	var ans = []byte{}
	for _, char := range []byte(s) {
		if char == 'i' {
			i := len(ans)
			j := i - 1
			for j >= i/2 {
				ans[j], ans[i-j-1] = ans[i-j-1], ans[j]
				j--
			}
		} else {
			ans = append(ans, char)
		}
	}
	return string(ans)
}

func revert(s string) string {
	fmt.Println(s)
	var str string
	for i := len(s) - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return str
}

func main() {
	fmt.Println(finalString("ksi"))
}
