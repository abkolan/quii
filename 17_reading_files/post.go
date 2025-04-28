package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	// instead of reading all the text, we need to scan line by line
	// so let's use bufio.Scanner

	scanner := bufio.NewScanner(postFile)

	readline := func(tagName string) string {
		scanner.Scan()
		line := scanner.Text()
		prefixTrimmed := strings.TrimPrefix(line, tagName)
		return prefixTrimmed
	}

	// Extract only the text from the separators
	title := readline(titleSeparator)
	description := readline(descriptionSeparator)
	tags := strings.Split(readline(tagsSeparator), ", ")
	body := readBody(scanner)

	post := Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}

	return post, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // skip the `---` line

	// Read body till EOF
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	// Remove the trailing new line
	body := strings.TrimSuffix(buf.String(), "\n")
	return body
}
