package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("zsh_history backup tool!")
	_, exists := os.LookupEnv("GH_TOKEN")
	if !exists {
		log.Fatal("Github Token environment (GH_TOKEN) variable not set!")
	}

	fmt.Println("Reading zsh_history file...")
	now := time.Now()

	homedir, err_homedir := os.UserHomeDir()
	check(err_homedir)

	readFile, err1 := os.Open(fmt.Sprintf("%s/.zsh_history", homedir))
	check(err1)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fmt.Sprintf("`%s`\n", fileScanner.Text()))
	}
	readFile.Close()

	fmt.Println("Creating Gist..")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GH_TOKEN")},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	gist := &github.Gist{
		Description: github.String("zsh_history file backup"),
		Public:      github.Bool(false),
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(fmt.Sprintf("zsh_history_%s.md", now.Format("2006_01_02_15_04_05"))): {
				Content: github.String(strings.Join(fileLines, "\n")),
			},
		},
	}

	_, _, err := client.Gists.Create(ctx, gist)
	check(err)
	fmt.Println("Gist created successfully!!")
}
