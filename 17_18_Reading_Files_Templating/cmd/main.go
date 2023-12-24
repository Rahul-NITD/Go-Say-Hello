package main

import (
	blogposts "GoSayHello/17_18_Reading_Files_Templating"
	blogrenderer "GoSayHello/17_18_Reading_Files_Templating/BlogRenderer"
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
