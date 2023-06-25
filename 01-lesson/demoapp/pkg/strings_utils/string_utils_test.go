package strings_utils

import "testing"

func TestRev(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "first test",
			s:    "DANIK",
			want: "KINAD",
		},
		{
			name: "second test",
			s:    "АПР",
			want: "РПА",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rev(tt.s); got != tt.want {
				t.Errorf("Rev() = %v, want %v", got, tt.want)
			}
		})
	}
}
