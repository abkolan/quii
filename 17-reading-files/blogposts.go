package blogposts

import (
	"io/fs"
)

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dirEntries, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dirEntries {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {

		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	//Open postFile
	postFile, err := fileSystem.Open(fileName)

	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)

}
