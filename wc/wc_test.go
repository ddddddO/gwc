package wc

import (
	"testing"
)

func Test_genHeader(t *testing.T) {
	tests := []struct {
		name string
		opts Options
		want string
	}{
		{
			name: "success(default)",
			opts: Options{
				IsHeader:        true,
				IsBytes:         false,
				IsChars:         false,
				IsLines:         false,
				IsWords:         false,
				IsMaxLineLength: false,
			},
			want: "  Lines  Words  Bytes\n",
		},
	}

	for _, tt := range tests {
		tmp := genHeader(tt.opts)
		got := string(tmp)
		if got != tt.want {
			t.Errorf("want: %s, got: %s\n", tt.want, got)
		}
	}
}
