package blogrenderer_test

import (
	"bytes"
	"testing"

	blogrenderer "github.com/abkolan/quii/18_templating"
)

func TestRender(t *testing.T) {
	var (
		post = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, post)

		if err != nil {
			t.Fatal("did not expect an error but got an error", err)
		}

		got := buf.String()
		want := `<h1>hello world</h1><p>This is a description</p>Tags: <ul><li>go</li><li>tdd</li></ul>`

		if got != want {
			t.Errorf("got `%s` but wanted `%s`", got, want)
		}
	})

}
