package main

import (
	"log"

	"github.com/divergentluna/git-prometheus-exporter/internal/collector"
	"github.com/divergentluna/git-prometheus-exporter/internal/config"
	"github.com/divergentluna/git-prometheus-exporter/internal/server"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	configPath := "configs/github_config.yaml"
	githubConfig, err := config.LoadGitHubConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	githubCollector := collector.NewGitHubCollector(githubConfig.GitHubOrg, githubConfig.GitHubRepo, githubConfig.GitHubToken)
	prometheus.MustRegister(githubCollector)

	server.StartServer(promhttp.Handler(), ":9101")
}
