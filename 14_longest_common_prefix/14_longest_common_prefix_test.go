package longestcommonprefix

import "testing"

func TestLongestPrefix(t *testing.T) {
	tests := []struct {
		Name  string
		Input []string
		Wait  string
	}{
		{"Simple 1", []string{"flower", "flow", "flight"}, "fl"},
		{"Simple 2", []string{"dog", "racecar", "car"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := longestCommonPrefix(tt.Input)
			if got != tt.Wait {
				t.Errorf("Got %s wait %s", got, tt.Wait)
			}
		})
	}
}

func BenchmarkLongestPrefix(b *testing.B) {
	strs := []string{"flower", "flow", "flight"}
	for i := 0; i < b.N; i++ {
		_ = longestCommonPrefix(strs)
	}
}
