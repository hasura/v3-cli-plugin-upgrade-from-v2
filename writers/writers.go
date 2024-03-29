package writers

// This provides a writer that acts as a proxy
// but removes duplicate sequential blank lines
// It changes from original buffering to character buffering

import (
	"fmt"
	"io"
	"regexp"
)

type RemoveDuplicateBlankLinesWriter struct {
	writer io.Writer
	regex  *regexp.Regexp
	blanks int
}

func NewRemoveDuplicateBlankLinesWriter(writer io.Writer) *RemoveDuplicateBlankLinesWriter {
	return &RemoveDuplicateBlankLinesWriter{
		writer: writer,
	}
}

// This just writes one byte at a time, but that's ok for now
// TODO: Optimise to build a buffer then write it all at once
func (w *RemoveDuplicateBlankLinesWriter) Write(p []byte) (n int, err error) {

	var count = 0

	for i := 0; i < len(p); i++ {
		b := p[i]
		s := fmt.Sprintf("%c", b)
		if b == '\n' {
			w.blanks++
			if w.blanks < 3 {
				c, e := w.writer.Write([]byte(s))
				count += c
				if e != nil {
					return count, e
				}
			}
		} else {
			if b != ' ' && b != '\t' {
				w.blanks = 0
			}
			c, e := w.writer.Write([]byte(s))
			count += c
			if e != nil {
				return count, e
			}
		}
	}

	return count, nil
}
