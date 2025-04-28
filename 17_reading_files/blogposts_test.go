package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/abkolan/quii/17_reading_files"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no always error")
}

func assertPost(t testing.TB, got blogposts.Post, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v but wanted %+v", got, want)
	}
}

func TestNewBlogPosts(t *testing.T) {

	t.Run("Returns an error", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})
		if err == nil {
			t.Errorf("Expected Error but did not get one")
		}

	})

	t.Run("Test values of Posts", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
			secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
Jai
Shri
Ram`
		)
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		got := posts[0]
		want := blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		}

		if err != nil {
			t.Error("got an error while wasn't expecting one")
		}

		assertPost(t, got, want)
	})

}
