package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"Simple case", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Another simple case", []int{3, 2, 4}, 6, []int{1, 2}},
		{"Duplicates", []int{3, 3}, 6, []int{0, 1}},
		{"Negative numbers", []int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
		{"Target is zero", []int{-3, 4, 3, 90}, 0, []int{0, 2}},
		{"Zeroes in array", []int{0, 4, 3, 0}, 0, []int{0, 3}},
		{"Large numbers", []int{1000000, 2000000, 3000000}, 5000000, []int{1, 2}},
		{"No solution", []int{1, 2, 3}, 7, []int{0, 0}}, // assuming default fallback
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FAIL: TwoSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15, 3, 6, 8, 5, 9, 10, 1, 4, 12, 14, 13}
	target := 9

	b.ResetTimer() // Optional: reset timer if setup before

	for i := 0; i < b.N; i++ {
		_ = twoSum(nums, target)
	}
}
