package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var (
	files       = make(map[github.GistFilename]github.GistFile)
	description string
	token       string
	secret      bool
)

var cmdGistup = &cobra.Command{
	Use:   "gistup",
	Short: "gistup uploads files to Github as gists",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		for _, f := range args {
			data, err := ioutil.ReadFile(f)
			if err != nil {
				log.Fatalf("gistup: reading file (%v)", err)
			}

			gf := github.GistFile{
				Content: github.String(string(data)),
			}

			name := filepath.Base(f)
			files[github.GistFilename(name)] = gf
		}

		gist := &github.Gist{
			Description: github.String(description),
			Files:       files,
			Public:      github.Bool(!secret),
		}

		oauthClient := oauth2.NewClient(oauth2.NoContext, &tokenSource{})
		client := github.NewClient(oauthClient)

		g, _, err := client.Gists.Create(gist)
		if err != nil {
			log.Fatalf("gistup: creating gist (%v)", err)
		}
		fmt.Println(*g.HTMLURL)
	},
}

// need this to create a client
type tokenSource struct{}

func (t *tokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{AccessToken: token}, nil
}

func init() {
	cmdGistup.Flags().StringVarP(&description, "description", "d", "", "description of gist")
	cmdGistup.Flags().StringVarP(&token, "token", "t", "", "use this github token")
	cmdGistup.Flags().BoolVarP(&secret, "secret", "s", false, "gist set to secret")
}

func main() {
	if token == "" {
		// load it from ~/.github
		path := filepath.Join(os.Getenv("HOME"), ".github")
		data, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("gistup: invalid token (%v)", err)
		}
		token = string(data)
	}

	cmdGistup.Execute()
}
