package blogrenderer

import (
	blogposts "GoSayHello/17_18_Reading_Files_Templating"
	"embed"
	"html/template"
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

const HTMLtemplate = `<h1>{{.Title}}</h1>
<p>{{.Description}}</p>
Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	postTemplate *template.Template
	mdParser     *parser.Parser
}

func NewPostRenderer() (PostRenderer, error) {
	templateToParseWith, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return PostRenderer{}, err
	}
	extensions := parser.AutoHeadingIDs | parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)
	return PostRenderer{templateToParseWith, p}, nil
}

func (p *PostRenderer) Render(w io.Writer, post blogposts.Post) error {
	return p.postTemplate.ExecuteTemplate(w, "blog.gohtml", p.GenerateBodyVM(post))
}

func (p *PostRenderer) RenderIndex(w io.Writer, posts []blogposts.Post) error {
	return p.postTemplate.ExecuteTemplate(w, "index.gohtml", posts)
}

type PostViewModel struct {
	Post     blogposts.Post
	HTMLBody template.HTML
}

func (p *PostRenderer) GenerateBodyVM(post blogposts.Post) PostViewModel {
	vm := PostViewModel{Post: post}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(post.Body), p.mdParser, nil))
	return vm
}
