package blogposts_test

import (
	blogposts "GoSayHello/BlogPosts"
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFileSystem struct{}

func (s StubFileSystem) Open(name string) (fs.File, error) {
	return nil, errors.New("I always fail")
}

func TestNewBlogPost(t *testing.T) {
	t.Run("Initial Test", func(t *testing.T) {
		blogContent := `Title: Post 1`
		fs := fstest.MapFS{
			"hello world.md": {Data: []byte(blogContent)},
		}
		blogPosts, _ := blogposts.NewBlogPosts(fs)
		if len(blogPosts) != len(fs) {
			t.Errorf("got %d want %d", len(blogPosts), len(fs))
		}
	})
	t.Run("Test if error occurs", func(t *testing.T) {
		fs := StubFileSystem{}
		_, err := blogposts.NewBlogPosts(fs)
		if err == nil {
			t.Fatal("Expected an Error")
		}
		if err.Error() != "I always fail" {
			t.Fatal("Wrong Error")
		}
	})
	t.Run("Test Tags", func(t *testing.T) {
		blogContent := `Title: Post 1
Description: This is post 1
Tags: firstPost, tmkb`
		fs := fstest.MapFS{
			"hello world.md": {Data: []byte(blogContent)},
		}
		bp, err := blogposts.NewBlogPosts(fs)
		AssertNoError(t, err)
		got := bp[0]
		want := blogposts.Post{Title: "Post 1", Description: "This is post 1", Tags: []string{"firstPost", "tmkb"}, Body: ""}
		AssertPosts(t, got, want)
	})
	t.Run("Test Body", func(t *testing.T) {
		blogContent := `Title: Post 1
Description: This is post 1
Tags: firstPost, tmkb
---
Hello World!
This is Body`
		fs := fstest.MapFS{
			"hello world.md": {Data: []byte(blogContent)},
		}
		bp, err := blogposts.NewBlogPosts(fs)
		AssertNoError(t, err)
		got := bp[0]
		want := blogposts.Post{
			Title:       "Post 1",
			Description: "This is post 1",
			Tags:        []string{"firstPost", "tmkb"},
			Body: `Hello World!
This is Body`,
		}
		AssertPosts(t, got, want)
	})
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Did not expect Error got %v", err.Error())
	}
}

func AssertPosts(t testing.TB, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}
