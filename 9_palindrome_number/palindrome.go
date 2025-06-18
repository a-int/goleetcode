package palindrome

func isPalindrome(x int) bool {
	original := x
	reverse := 0
	for x > 0 {
		reverse = (reverse * 10) + x%10
		x /= 10
	}
	return (original == reverse)
}
