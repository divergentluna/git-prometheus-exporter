package config

type GitHubExporterConfig struct {
	GitHubToken string `yaml:"github_token"`
	GitHubRepo  string `yaml:"github_repo"`
	GitHubOrg   string `yaml:"github_org"`
	MetricsPort int    `yaml:"metrics_port"`
}
