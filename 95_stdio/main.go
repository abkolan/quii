/*
Read stdio in 512 byte chunks and print it on stdout
*/

package main

import (
	"io"
	"os"
)

func main() {
	buf := make([]byte, 512)
	for {
		n, err := os.Stdin.Read(buf)
		if n > 0 {
			if _, werr := os.Stdout.Write(buf[:n]); werr != nil {
				os.Exit(1)
			}
		}
		if err != nil {
			if err == io.EOF {
				// clean
				break
			}
			os.Exit(1)
		}
	}
}
