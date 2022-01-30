package main

import "fmt"

func RomanNumeralConversion(s string) (ans int) {
	mp := make(map[byte]int)
	mp['M'] = 1000
	mp['D'] = 500
	mp['C'] = 100
	mp['L'] = 50
	mp['X'] = 10
	mp['V'] = 5
	mp['I'] = 1

	var before byte
	for i := 0; i < len(s); i++ {
		if i > 0 {
			before = s[i-1]
		}
		now := mp[s[i]]
		if now == 0 {
			return -1
		} else {
			if mp[before] > now {
				ans -= now
			} else {
				ans += now
			}
		}
	}
	return
}

func main() {
	ans := RomanNumeralConversion("MCM")
	if ans == -1 {
		fmt.Println("Invalid conversion")
	} else {
		fmt.Println(ans)
	}
}
