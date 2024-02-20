package writers

import (
	"io"
	"regexp"
)

type RemoveDuplicateBlankLinesWriter struct {
	writer io.Writer
	regex  *regexp.Regexp
}

func NewRemoveDuplicateBlankLinesWriter(writer io.Writer) *RemoveDuplicateBlankLinesWriter {
	// Define a regular expression pattern to match consecutive blank lines
	regexPattern := `\n\s*\n+`

	// Compile the regular expression
	regex := regexp.MustCompile(regexPattern)

	return &RemoveDuplicateBlankLinesWriter{
		writer: writer,
		regex:  regex,
	}
}

func (w *RemoveDuplicateBlankLinesWriter) Write(p []byte) (n int, err error) {
	// Convert the byte slice to a string
	inputString := string(p)

	// Replace consecutive blank lines with a single blank line
	// Warning: This may be split over half the blank lines
	resultString := w.regex.ReplaceAllString(inputString, "\n\n")

	// Convert the modified string back to a byte slice
	resultBytes := []byte(resultString)

	// Write the modified byte slice to the underlying writer (os.Stdout in this case)
	return w.writer.Write(resultBytes)
}
