package upppercaser

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func UpperCase(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	firstLine := true
	for scanner.Scan() {
		if !firstLine {
			fmt.Fprintln(w)
		} else {
			firstLine = false
		}
		line := scanner.Text()
		_, err := fmt.Fprint(w, strings.ToUpper(line))
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}

func CountBytes(r io.Reader) (int64, error) {
	return io.Copy(io.Discard, r)
}
