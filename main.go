package main

import (
	"log"
	"os"
)

func main() {
	bs, err := os.ReadFile(os.Getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		log.Fatalf("failed to read github event: %v", err)
	}
	log.Println(string(bs))
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
