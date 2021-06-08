package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

type GithubEvent struct {
	Issue struct {
		Number int
	}
	Repository struct {
		FullName string `json:"full_name"`
	}
}

func main() {
	bs, err := os.ReadFile(os.Getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		log.Fatalf("failed to read github event: %v", err)
	}
	var event GithubEvent
	if err := json.Unmarshal(bs, &event); err != nil {
		log.Fatalf("failed to unmarshal event: %v", err)
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	names := strings.Split(event.Repository.FullName, "/")
	owner := names[0]
	repo := names[1]
	pr, _, err := client.PullRequests.Get(ctx, owner, repo, event.Issue.Number)
	if err != nil {
		log.Fatalf("failed to get pull requests: %v", err)
	}
	fmt.Println(pr.Head.Ref)
	runs, _, err := client.Checks.ListCheckRunsForRef(ctx, owner, repo, *pr.Head.Ref, nil)
	if err != nil {
		log.Fatalf("failed to get check runs: %v", err)
	}
	toRerun := map[int64]bool{}
	for _, run := range runs.CheckRuns {
		if *run.Conclusion == "failure" || *run.Conclusion == "cancelled" {
			toRerun[*run.CheckSuite.ID] = true
		}
	}

	for suite := range toRerun {
		if _, err := client.Checks.ReRequestCheckSuite(ctx, owner, repo, suite); err != nil {
			log.Fatalf("failed to rerun suite: %v", err)
		}
	}
}
