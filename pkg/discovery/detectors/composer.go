package detectors

import (
	"github.com/pvaass/repo-analyzer/pkg/repository"
)

type Composer struct{}

func (d Composer) Detect(repo repository.Repository, resultChannel chan Result) {
	result := Result{
		Identifier: "composer",
	}
	if hasComposer(repo) {
		result.Score = 100
	}
	resultChannel <- result
}

func hasComposer(repo repository.Repository) bool {
	for _, name := range repo.FileNames() {
		if name == "composer.json" {
			return true
		}
	}

	return false
}

func init() {
	register(Composer{})
}
