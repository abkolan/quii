package upppercaser_test

import (
	"bytes"
	"strings"
	"testing"

	upppercaser "github.com/abkolan/quii/96_io"
)

func TestUpperCaser(t *testing.T) {
	cases := []struct {
		Description string
		Input       string
		Output      string
	}{
		{
			"One word",
			"hello",
			"HELLO",
		},
		{
			"With New line",
			`Hello
World`,
			`HELLO
WORLD`,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			reader := strings.NewReader(c.Input)
			buf := bytes.Buffer{}

			err := upppercaser.UpperCase(reader, &buf)

			if err != nil {
				t.Errorf("got error when it wasn't expected %v", err)
			}
			want := c.Output
			got := buf.String()

			if got != want {
				t.Errorf("got %q but want %q", got, want)
			}

		})
	}

}

func TestCountBytes(t *testing.T) {

	lengthFunc := func(s string) int64 {
		b := []byte(s)
		return int64(len(b))
	}

	cases := []struct {
		Description string
		Input       string
	}{
		{
			Description: "one word",
			Input:       "hello",
		},
		{
			Description: "Empty",
			Input:       "",
		},
		{
			Description: "Unicode",
			Input:       "अ आ इ ई उ ऊ ऋ ऌ ऍ ऎ ए ऐ ऑ ऒ ओ औ क ख ग घ ङ",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Description, func(t *testing.T) {
			want := lengthFunc(c.Input)
			got, err := upppercaser.CountBytes(strings.NewReader(c.Input))

			if err != nil {
				t.Fatalf("got an error %v when it wasn't expected", err)
			}

			if got != want {
				t.Errorf("got %q but wanted %q", got, want)
			}
		})
	}
}
