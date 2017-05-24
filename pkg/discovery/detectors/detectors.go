package detectors

import (
	"github.com/pvaass/repo-analyzer/pkg/repository"
)

var collection []Detector

type Detector interface {
	Identifier() string
	Detect(repo repository.Repository) int
}

type Result struct {
	Identifier string
	Score      int
}

func Run(repo repository.Repository) []Result {
	var results []Result
	for _, detector := range collection {
		results = append(results, Result{detector.Identifier(), detector.Detect(repo)})
	}
	return results
}

func register(detector Detector) {
	collection = append(collection, detector)
}
