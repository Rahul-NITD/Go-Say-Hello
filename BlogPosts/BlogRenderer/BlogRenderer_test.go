package blogrenderer_test

import (
	blogposts "GoSayHello/BlogPosts"
	blogrenderer "GoSayHello/BlogPosts/BlogRenderer"
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestBlogRenderer(t *testing.T) {
	aPost := blogposts.Post{
		Title:       "Hello Rahul",
		Description: "This is some description",
		Tags:        []string{"TDD", "something else too"},
		Body:        "Ooh lala",
	}
	t.Run("convert single post into HTML", func(t *testing.T) {
		b := bytes.Buffer{}
		postrenderer, err := blogrenderer.NewPostRenderer()
		if err != nil {
			t.Fatal(err.Error())
		}
		postrenderer.Render(&b, aPost)
		approvals.VerifyString(t, b.String())
	})

	t.Run("Render an index page", func(t *testing.T) {
		b := bytes.Buffer{}
		postrenderer, err := blogrenderer.NewPostRenderer()
		if err != nil {
			t.Fatal(err.Error())
		}
		posts := []blogposts.Post{
			{Title: "Hello World"},
			{Title: "Hello Again"},
		}
		if err := postrenderer.RenderIndex(&b, posts); err != nil {
			t.Fatal(err.Error())
		}
		approvals.VerifyString(t, b.String())
	})
}
