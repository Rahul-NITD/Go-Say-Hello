package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func newPost(file io.Reader) (Post, error) {

	scanner := bufio.NewScanner(file)
	readMetaLine := func(metaTag string) string {
		scanner.Scan()
		if len(scanner.Text()) > 0 {
			return strings.TrimPrefix(scanner.Text(), metaTag)
		}
		return ""
	}
	readBody := func() string {
		scanner.Scan()
		buf := bytes.Buffer{}
		for scanner.Scan() {
			fmt.Fprintln(&buf, scanner.Text())
		}
		return strings.TrimSuffix(buf.String(), "\n")
	}

	return Post{
		Title:       readMetaLine("Title: "),
		Description: readMetaLine("Description: "),
		Tags:        strings.Split(readMetaLine("Tags: "), ", "),
		Body:        readBody(),
	}, nil

}

func (p *Post) SanitizeTitle() string {
	return strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
}
