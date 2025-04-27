package blogposts

import (
	"io/fs"
	"testing/fstest"
)

type Post struct {
}

func NewPostsFromFS(fileSystem fstest.MapFS) []Post {
	dirEntries, _ := fs.ReadDir(fileSystem, ".")
	var posts []Post
	for range dirEntries {
		posts = append(posts, Post{})
	}
	return posts
}
