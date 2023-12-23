package blogposts

import (
	"io/fs"
)

func NewBlogPosts(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err // Do we want to stop if one file fails?
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	// open file
	file, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, nil
	}
	defer file.Close()

	return newPost(file)
}
