package main

import (
	blogposts "GoSayHello/BlogPosts"
	blogrenderer "GoSayHello/BlogPosts/BlogRenderer"
	"log"
	"os"
)

func main() {
	posts, err := blogposts.NewBlogPosts(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	postrenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		log.Fatal(err)
	}
	postrenderer.Render(os.Stdout, posts[0])
}
