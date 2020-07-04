package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	start, max := 0, 0
	// makes a slice (expandable) which is type map
	// with key[rune]:value[int]
	mp := make(map[rune]int)
	for i, c := range s {
		if prevIndex, ok := mp[c]; ok && mp[c] >= start {
			start = prevIndex + 1
			mp[c] = i
			continue
		}
		if dist := i - start + 1; dist > max {
			max = dist
		}
		mp[c] = i
	}
	return max
}

func main() {
	input := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(input))
}
