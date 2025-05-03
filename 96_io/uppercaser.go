package upppercaser

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

func UpperCase(r io.Reader, w io.Writer) error {
	br := bufio.NewReader(r)
	for {
		line, err := br.ReadString('\n')
		if len(line) > 0 {
			if _, werr := w.Write([]byte(strings.ToUpper(line))); werr != nil {
				return werr
			}
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				// clean
				break
			}
			return err
		}

	}
	return nil
}

func CountBytes(r io.Reader) (int64, error) {
	return io.Copy(io.Discard, r)
}
