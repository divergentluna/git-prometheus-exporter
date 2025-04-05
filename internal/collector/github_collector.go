package collector

import (
	"context"

	"github.com/divergentluna/git-prometheus-exporter/internal/client"
	"github.com/prometheus/client_golang/prometheus"
)

type GitHubCollector struct {
	repo   string
	org    string
	client *client.GitHubClient

	stars   *prometheus.Desc
	forks   *prometheus.Desc
	openPRs *prometheus.Desc

	commits          *prometheus.Desc
	mergedPRs        *prometheus.Desc
	openIssues       *prometheus.Desc
	totalReviews     *prometheus.Desc
	avgReviewsPerPR  *prometheus.Desc
	staleBranches    *prometheus.Desc
	activeBranches   *prometheus.Desc
	repoStars        *prometheus.Desc
	repoForks        *prometheus.Desc
}

func NewGitHubCollector(org, repo, token string) *GitHubCollector {
	return &GitHubCollector{
		repo:   repo,
		org:    org,
		client: client.NewGitHubClient(token),

		stars:   prometheus.NewDesc("github_repo_stars_total", "Number of stars", nil, nil),
		forks:   prometheus.NewDesc("github_repo_forks_total", "Number of forks", nil, nil),
		openPRs: prometheus.NewDesc("github_open_pull_requests_total", "Number of open PRs", nil, nil),

		commits:         prometheus.NewDesc("git_commits_total", "Total number of commits", nil, nil),
		mergedPRs:       prometheus.NewDesc("git_prs_merged_total", "Total number of merged pull requests", nil, nil),
		openIssues:      prometheus.NewDesc("git_issues_open", "Number of open issues", nil, nil),
		totalReviews:    prometheus.NewDesc("git_reviews_total", "Total number of reviews", nil, nil),
		avgReviewsPerPR: prometheus.NewDesc("git_avg_reviews_per_pr", "Average number of reviews per pull request", nil, nil),
		staleBranches:   prometheus.NewDesc("git_stale_branches_total", "Number of stale branches", nil, nil),
		activeBranches:  prometheus.NewDesc("git_active_branches_total", "Number of active branches", nil, nil),
		repoStars:       prometheus.NewDesc("git_repo_stars", "Number of stars", nil, nil),
		repoForks:       prometheus.NewDesc("git_repo_forks", "Number of forks", nil, nil),
	}
}

func (c *GitHubCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.stars
	ch <- c.forks
	ch <- c.openPRs
	ch <- c.commits
	ch <- c.mergedPRs
	ch <- c.openIssues
	ch <- c.totalReviews
	ch <- c.avgReviewsPerPR
	ch <- c.staleBranches
	ch <- c.activeBranches
	ch <- c.repoStars
	ch <- c.repoForks
}

func (c *GitHubCollector) Collect(ch chan<- prometheus.Metric) {
	ctx := context.Background()
	stats, err := c.client.FetchRepoStats(ctx, c.org, c.repo)
	if err != nil {
		return
	}

	ch <- prometheus.MustNewConstMetric(c.stars, prometheus.GaugeValue, float64(stats.Stars))
	ch <- prometheus.MustNewConstMetric(c.forks, prometheus.GaugeValue, float64(stats.Forks))
	ch <- prometheus.MustNewConstMetric(c.openPRs, prometheus.GaugeValue, float64(stats.OpenPRs))
	ch <- prometheus.MustNewConstMetric(c.commits, prometheus.CounterValue, float64(stats.Commits))
	ch <- prometheus.MustNewConstMetric(c.mergedPRs, prometheus.CounterValue, float64(stats.MergedPRs))
	ch <- prometheus.MustNewConstMetric(c.openIssues, prometheus.GaugeValue, float64(stats.OpenIssues))
	ch <- prometheus.MustNewConstMetric(c.totalReviews, prometheus.CounterValue, float64(stats.TotalReviews))
	ch <- prometheus.MustNewConstMetric(c.avgReviewsPerPR, prometheus.GaugeValue, stats.AvgReviewsPerPR)
	ch <- prometheus.MustNewConstMetric(c.staleBranches, prometheus.GaugeValue, float64(stats.StaleBranches))
	ch <- prometheus.MustNewConstMetric(c.activeBranches, prometheus.GaugeValue, float64(stats.ActiveBranches))
	ch <- prometheus.MustNewConstMetric(c.repoStars, prometheus.GaugeValue, float64(stats.Stars))
	ch <- prometheus.MustNewConstMetric(c.repoForks, prometheus.GaugeValue, float64(stats.Forks))
}
