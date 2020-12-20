package wc

import (
	"bufio"
	"bytes"

	"fmt"

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

type fileInfo struct {
	name    string
	lineCnt int
	wordCnt int
	charCnt int
	byteCnt int
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

	fis := []*fileInfo{}
	for _, file := range files {
		fi := &fileInfo{}
		if err := fi.wc(file); err != nil {
			return err
		}
		fis = append(fis, fi)
	}

	for _, fi := range fis {
		b := []byte{}
		if !opts.IsBytes && !opts.IsChars && !opts.IsLines {
			defaultTemplate := "  %d  %d  %d  %s\n"
			b = append(b, []byte(fmt.Sprintf(defaultTemplate, fi.lineCnt, fi.wordCnt, fi.byteCnt, fi.name))...)
			_, err := buf.Write(b)
			if err != nil {
				return err
			}
			continue
		}

		template := "  %d"
		if opts.IsLines {
			b = append(b, []byte(fmt.Sprintf(template, fi.lineCnt))...)
		}
		if opts.IsWords {
			b = append(b, []byte(fmt.Sprintf(template, fi.wordCnt))...)
		}
		if opts.IsChars {
			b = append(b, []byte(fmt.Sprintf(template, fi.charCnt))...)
		}
		if opts.IsBytes {
			b = append(b, []byte(fmt.Sprintf(template, fi.byteCnt))...)
		}
		b = append(b, []byte(fmt.Sprintf("  %s", fi.name))...)
		b = append(b, []byte("\n")...)

		_, err := buf.Write(b)
		if err != nil {
			return err
		}
	}

	io.Copy(os.Stdout, buf)
	return nil
}

func (fi *fileInfo) wc(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	fi.name = f.Name()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fi.lineCnt++
		fi.charCnt += 2 // 改行分
		fi.byteCnt += 2 // 改行分

		line := scanner.Text()

		word := ""
		for i, char := range line {
			fi.charCnt++
			fi.byteCnt += len(string(char))

			if len(word) > 0 && (string(char) == " ") {
				fi.wordCnt++
				word = ""
				continue
			}

			if string(char) == " " {
				word = ""
				continue
			}

			if len(word) > 0 && (i == len(line)-1) {
				fi.wordCnt++
			}

			word += string(char)
		}
	}
	// ...
	fi.lineCnt--
	fi.charCnt -= 2
	fi.byteCnt -= 2

	return nil
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
	if opts.IsChars {
		buf = append(buf, []byte("  Chars")...)
	}
	if opts.IsBytes {
		buf = append(buf, []byte("  Bytes")...)
	}
	buf = append(buf, []byte("\n")...)

	return buf
}
