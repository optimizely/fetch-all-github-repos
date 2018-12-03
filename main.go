package main

import (
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"os"
	"os/exec"
	"strings"
)

func runCommand(command string, ignoreErrors bool) {
	fmt.Println(command)
	out, err := exec.Command("bash", "-c", command).Output()
	if !ignoreErrors && err != nil {
		fmt.Printf("Error running command: %s", err)
	}
	fmt.Printf("%s", out)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("You must supply an access_token\n")
		os.Exit(1)
	}
	access_token := os.Args[1]
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: access_token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := &github.RepositoryListByOrgOptions{
		Type:        "all",
		ListOptions: github.ListOptions{PerPage: 10},
	}

	// get all pages of results
	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, "optimizely", opt)
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	/*
		for _, repo := range allRepos {
			// change https:// to git@
			gitRepo := strings.Replace(*repo.CloneURL, "https://github.com/", "git@github.com:", 1)
			fmt.Printf("%s\n", gitRepo)
		}
	*/

	// now clone each and search it
	for j, repo := range allRepos {
		// change https:// to git@
		gitRepo := strings.Replace(*repo.CloneURL, "https://github.com/", "git@github.com:", 1)
		fmt.Printf("repo number: %d - %s\n", j, gitRepo)
		command := "git clone --depth 1 " + gitRepo + " tmpdir"
		runCommand(command, false)
		runCommand("ag flatmap-stream tmpdir", true)
		runCommand("ag package.json tmpdir", true)
		runCommand("rm -rf tmpdir", false)
	}
}
