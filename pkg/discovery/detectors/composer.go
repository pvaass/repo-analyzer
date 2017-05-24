package detectors

import (
	"github.com/pvaass/repo-analyzer/pkg/repository"
)

type Composer struct{}

func (d Composer) Identifier() string {
	return "composer"
}

func (d Composer) Detect(repo repository.Repository) int {
	for _, name := range repo.FileNames() {
		if name == "composer.json" {
			return 100
		}
	}

	return 0
}
