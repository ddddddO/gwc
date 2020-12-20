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
		{
			name: "success(Bytes)",
			opts: Options{
				IsHeader:        true,
				IsBytes:         true,
				IsChars:         false,
				IsLines:         false,
				IsWords:         false,
				IsMaxLineLength: false,
			},
			want: "  Bytes\n",
		},
		{
			name: "success(Chars)",
			opts: Options{
				IsHeader:        true,
				IsBytes:         false,
				IsChars:         true,
				IsLines:         false,
				IsWords:         false,
				IsMaxLineLength: false,
			},
			want: "  Chars\n",
		},
		{
			name: "success(Lines)",
			opts: Options{
				IsHeader:        true,
				IsBytes:         false,
				IsChars:         false,
				IsLines:         true,
				IsWords:         false,
				IsMaxLineLength: false,
			},
			want: "  Lines\n",
		},
		{
			name: "success(Bytes, Lines)",
			opts: Options{
				IsHeader:        true,
				IsBytes:         true,
				IsChars:         false,
				IsLines:         true,
				IsWords:         false,
				IsMaxLineLength: false,
			},
			want: "  Lines  Bytes\n",
		},
		{
			name: "success(Bytes, Chars, Lines)",
			opts: Options{
				IsHeader:        true,
				IsBytes:         true,
				IsChars:         true,
				IsLines:         true,
				IsWords:         false,
				IsMaxLineLength: false,
			},
			want: "  Lines  Chars  Bytes\n",
		},
		{
			name: "success(Lines, Words, Bytes)",
			opts: Options{
				IsHeader:        true,
				IsBytes:         true,
				IsChars:         false,
				IsLines:         true,
				IsWords:         true,
				IsMaxLineLength: false,
			},
			want: "  Lines  Words  Bytes\n",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tmp := genHeader(tt.opts)
			got := string(tmp)
			if got != tt.want {
				t.Errorf("\nwant: %sgot: %s\n", tt.want, got)
			}
		})

	}
}
