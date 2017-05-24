package detectors

import (
	"github.com/pvaass/repo-analyzer/pkg/repository"
)

var collection []Detector

type Detector interface {
	Detect(repo repository.Repository, resultChannel chan Result)
}

type Result struct {
	Identifier string
	Score      int
}

func Run(repo repository.Repository) []Result {

	if hasComposer(repo) {
		repo.File("composer.json")
	}

	if hasNpm(repo) {
		repo.File("package.json")
	}

	resultChannel := make(chan Result)
	for _, detector := range collection {
		go detector.Detect(repo, resultChannel)
	}
	var results []Result
	for i := 0; i < len(collection); i++ {
		results = append(results, <-resultChannel)
	}
	return results
}

func register(detector Detector) {
	collection = append(collection, detector)
}
