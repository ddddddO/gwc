package wc

import (
	"bytes"

	"io"
	"os"
)

// Options is ...
type Options struct {
	IsHeader bool // original option!

	IsBytes         bool // wc -c
	IsChars         bool // wc -m
	IsLines         bool // wc -l
	IsWords         bool // wc -w
	IsMaxLineLength bool // wc -L
}

// Wc is ...
func Wc(opts Options, files []string) error {
	buf := &bytes.Buffer{}

	if opts.IsHeader {
		header := genHeader(opts)
		_, err := buf.Write(header)
		if err != nil {
			return err
		}
	}

	io.Copy(os.Stdout, buf)
	return nil
}

func wc(opts Options, file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return nil, nil
}

func genHeader(opts Options) []byte {
	if !opts.IsHeader {
		return nil
	}

	if !opts.IsBytes && !opts.IsChars && !opts.IsLines {
		return []byte("  Lines  Words  Bytes\n")
	}

	buf := []byte{}
	if opts.IsLines {
		buf = append(buf, []byte("  Lines")...)
	}
	if opts.IsWords {
		buf = append(buf, []byte("  Words")...)
	}
	if opts.IsBytes {
		buf = append(buf, []byte("  Bytes")...)
	}
	buf = append(buf, []byte("\n")...)

	return buf
}
