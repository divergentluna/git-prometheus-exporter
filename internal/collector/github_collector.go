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
}

func NewGitHubCollector(org, repo, token string) *GitHubCollector {
	return &GitHubCollector{
		repo:   repo,
		org:    org,
		client: client.NewGitHubClient(token),

		stars:   prometheus.NewDesc("github_repo_stars_total", "Number of stars", nil, nil),
		forks:   prometheus.NewDesc("github_repo_forks_total", "Number of forks", nil, nil),
		openPRs: prometheus.NewDesc("github_open_pull_requests_total", "Number of open PRs", nil, nil),
	}
}

func (c *GitHubCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.stars
	ch <- c.forks
	ch <- c.openPRs
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
}
