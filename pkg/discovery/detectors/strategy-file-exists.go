package detectors

import (
	"github.com/pvaass/repo-analyzer/pkg/repository"
)

type FileExistsDetector struct{}

func (d FileExistsDetector) Init(repo repository.Repository) {}

func (d FileExistsDetector) Supports(rule Rule) bool {
	return rule.Strategy == "file-exist"
}

func (d FileExistsDetector) Detect(repo repository.Repository, resultChannel chan Result, rule Rule) {
	result := Result{
		Identifier: rule.Name,
	}
	if hasFile(repo, rule) {
		result.Score = 100
	}
	resultChannel <- result
}

func hasFile(repo repository.Repository, rule Rule) bool {
	return len(getFile(repo, rule)) > 0
}

func getFile(repo repository.Repository, rule Rule) []byte {
	return findFile(
		repo,
		rule.Arguments,
	)
}

func init() {
	register(&FileExistsDetector{})
}
