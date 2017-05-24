package analyze

import (
	"log"

	"github.com/pvaass/repo-analyzer/pkg/discovery/detectors"
	"github.com/pvaass/repo-analyzer/pkg/repository"
)

// Run analyzes a Repository for known languages and frameworks
func Run(repo repository.Repository) {
	detectors := [...]detector{
		detectors.Composer{},
		detectors.Symfony{},
		detectors.Laravel{},
	}

	var results []result
	for _, detector := range detectors {
		results = append(results, result{detector.Identifier(), detector.Detect(repo)})
	}

	log.Println(results)
}

type detector interface {
	Detect(repo repository.Repository) int
	Identifier() string
}
type result struct {
	identifier string
	certainty  int
}
