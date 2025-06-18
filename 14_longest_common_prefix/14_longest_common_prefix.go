package longestcommonprefix

import "strings"

func strLen(s string) int {
	return (len([]rune(s)))
}

func longestCommonPrefix(strs []string) string {
	minStr := strs[0]

	for _, curr := range strs {
		if strLen(curr) <= strLen(minStr) {
			minStr = curr
		}
	}

	stop := 1
	for stop <= strLen(minStr) {
		for _, str := range strs {
			if !strings.HasPrefix(str, minStr[:stop]) {
				return minStr[:stop-1]
			}
		}
		stop += 1
	}
	return ""
}
