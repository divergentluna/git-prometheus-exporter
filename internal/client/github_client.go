package client

import (
	"context"
	"time"

	"github.com/google/go-github/v55/github"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	client *github.Client
}

type RepoStats struct {
	Stars           int
	Forks           int
	OpenPRs         int
	Commits         int
	MergedPRs       int
	OpenIssues      int
	TotalReviews    int
	AvgReviewsPerPR float64
	StaleBranches   int
	ActiveBranches  int
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

	commits, _, err := g.client.Repositories.ListCommits(ctx, org, repo, nil)
	if err != nil {
		return nil, err
	}

	issues, _, err := g.client.Issues.ListByRepo(ctx, org, repo, &github.IssueListByRepoOptions{State: "open"})
	if err != nil {
		return nil, err
	}

	totalReviews := 0
	for _, pr := range prList {
		reviews, _, err := g.client.PullRequests.ListReviews(ctx, org, repo, pr.GetNumber(), nil)
		if err != nil {
			return nil, err
		}
		totalReviews += len(reviews)
	}

	avgReviewsPerPR := 0.0
	if len(prList) > 0 {
		avgReviewsPerPR = float64(totalReviews) / float64(len(prList))
	}

	branches, _, err := g.client.Repositories.ListBranches(ctx, org, repo, nil)
	if err != nil {
		return nil, err
	}

	staleBranches := 0
	activeBranches := 0
	for _, branch := range branches {
		branchInfo, _, err := g.client.Repositories.GetBranch(ctx, org, repo, branch.GetName(), false)
		if err != nil {
			return nil, err
		}
		if branchInfo.GetCommit().GetCommit().GetCommitter().GetDate().Before(time.Now().AddDate(0, -6, 0)) {
			staleBranches++
		} else {
			activeBranches++
		}
	}

	return &RepoStats{
		Stars:           repoData.GetStargazersCount(),
		Forks:           repoData.GetForksCount(),
		OpenPRs:         len(prList),
		Commits:         len(commits),
		MergedPRs:       0, // Placeholder, as mergedPRs is not calculated
		OpenIssues:      len(issues),
		TotalReviews:    totalReviews,
		AvgReviewsPerPR: avgReviewsPerPR,
		StaleBranches:   staleBranches,
		ActiveBranches:  activeBranches,
	}, nil
}
