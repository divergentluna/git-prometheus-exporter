package client

import (
	"context"

	"github.com/google/go-github/v55/github"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	client *github.Client
}

type RepoStats struct {
	Stars   int
	Forks   int
	OpenPRs int
}

func NewGitHubClient(token string) *GitHubClient {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.Background(), ts)

	return &GitHubClient{
		client: github.NewClient(tc),
	}
}

func (g *GitHubClient) FetchRepoStats(ctx context.Context, org, repo string) (*RepoStats, error) {
	repoData, _, err := g.client.Repositories.Get(ctx, org, repo)
	if err != nil {
		return nil, err
	}

	prList, _, err := g.client.PullRequests.List(ctx, org, repo, nil)
	if err != nil {
		return nil, err
	}

	return &RepoStats{
		Stars:   repoData.GetStargazersCount(),
		Forks:   repoData.GetForksCount(),
		OpenPRs: len(prList),
	}, nil
}
