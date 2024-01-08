package main

import (
	"fmt"
	"strings"
)

// Without using any string methods
//func WordCount(s string) map[string]int {
//	cur := ""
//	mp := make(map[string]int)
//	for i := 0; i < len(s); i++ {
//		if s[i] != ' ' {
//			cur = cur + string(s[i])
//		} else {
//			_, ok := mp[cur]
//			if ok {
//				mp[cur] += 1
//			} else {
//				mp[cur] = 1
//			}
//			cur = ""
//		}
//	}
//	_, ok := mp[cur]
//	if ok {
//		mp[cur] += 1
//	} else {
//		mp[cur] = 1
//	}
//	cur = ""
//	return mp
//}

func WordCount(s string) map[string]int {
	mp := make(map[string]int)
	for _, value := range strings.Fields(s) {
		_, ok := mp[value]
		if ok {
			mp[value] += 1
		} else {
			mp[value] = 1
		}
	}
	return mp
}

func main() {
	fmt.Println(WordCount("a a a a a bh b b b b "))
}
