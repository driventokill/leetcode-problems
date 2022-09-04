package bullsandcows

import "testing"

func Test_getHint(t *testing.T) {
	tests := []struct {
		name   string
		secret string
		guess  string
		want   string
	}{
		{"1", "8921", "2321", "2A0B"},
		{"2", "1807", "7810", "1A3B"},
		{"3", "1123", "0111", "1A1B"},
		{"4", "1122", "1222", "3A0B"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHint1(tt.secret, tt.guess); got != tt.want {
				t.Errorf("getHint() = %v, want %v", got, tt.want)
			}
		})
	}
}
