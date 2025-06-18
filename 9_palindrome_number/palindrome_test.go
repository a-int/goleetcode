package palindrome

import (
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	tests := []struct {
		Name  string
		Input int
		Wait  bool
	}{
		{"Simple 1", 121, true},
		{"Simple 2", -121, false},
		{"Simple 3", 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := isPalindrome(tt.Input)
			if got != tt.Wait {
				t.Errorf("Got %t instead of %t for input %d", got, tt.Wait, tt.Input)
			}
		})
	}
}

func BenchmarkPalindromeBigNumber(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isPalindrome(1919191991919191)
	}
}

func BenchmarkPalindromeSmallNumber(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isPalindrome(1221)
	}
}

func BenchmarkPalindromeWrongNumber(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isPalindrome(191919123919191)
	}
}
