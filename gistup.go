// CLI for quickly uploading gists.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

var (
	desc   = ""
	public = false
	files  gistFiles
	gists  = make(map[github.GistFilename]github.GistFile)
)

func init() {
	flag.StringVar(&desc, "description", desc, "description of the gist")
	flag.StringVar(&desc, "d", desc, "description of the gist")
	flag.BoolVar(&public, "public", public, "whether the file is public or not, defaults to private")
	flag.BoolVar(&public, "p", public, "whether the file is public or not, defaults to private")
	flag.Var(&files, "files", "files required for gist, space separated")
	flag.Var(&files, "f", "files required for gist, space separated")
}

type gistFiles []string

func (g *gistFiles) Set(value string) error {
	if len(*g) > 0 {
		return errors.New("files flag already set")
	}

	for _, f := range strings.Split(value, " ") {
		*g = append(*g, f)
	}
	return nil
}

func (g *gistFiles) String() string {
	return fmt.Sprint(*g)
}

// util for error handling
func exit(err error) {
	fmt.Printf("ERROR: %s\n", err)
	os.Exit(1)
}

func main() {
	flag.Parse()
	if len(files) == 0 || desc == "" {
		fmt.Printf("Missing description and/or files\n")
		flag.Usage()
		return
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Printf("Set your GITHUB_TOKEN variable!\n")
		return
	}

	for _, f := range files {
		bytes, err := ioutil.ReadFile(f)
		if err != nil {
			exit(err)
		}
		s := string(bytes)
		gist := github.GistFile{
			Content:  github.String(s),
			Filename: github.String(f),
		}
		gists[github.GistFilename(f)] = gist
	}
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}
	client := github.NewClient(t.Client())
	gist := &github.Gist{
		Description: &desc,
		Files:       gists,
		Public:      github.Bool(public),
	}
	g, _, err := client.Gists.Create(gist)
	if err != nil {
		exit(err)
	}
	fmt.Printf("Successfully uploaded Gist ... Public: %t Description: \"%s\"\n", *g.Public, *g.Description)
}
