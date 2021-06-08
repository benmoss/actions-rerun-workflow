package main

import (
	"encoding/json"
	"log"
	"os"
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
	log.Println(string(bs))
	var event GithubEvent
	if err := json.Unmarshal(bs, &event); err != nil {
		log.Fatalf("failed to unmarshal event: %v", err)
	}
	log.Printf("%#v", event)
	// ctx := context.Background()
	// ts := oauth2.StaticTokenSource{
	// 	&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	// }
	// tc := oauth2.NewClient(ctx, ts)
	// client := github.NewClient(tc)
	// client.PullRequests.Get(context.TODO(), os.Getenv("REPO_OWNER"), os.Getenv("REPO_NAME"), os.Getenv(")
	// orgs, _, err := client.Organizations.List(context.Background(), "willnorris", nil)
	// fmt.Println(orgs, err)
}
