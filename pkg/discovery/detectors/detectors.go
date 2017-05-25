package detectors

import (
	"github.com/pvaass/repo-analyzer/pkg/repository"
)

var collection []Detector

type Detector interface {
	Detect(repo repository.Repository, resultChannel chan Result, rule Rule)
	Supports(rule Rule) bool
	Init(repo repository.Repository)
}

type Result struct {
	Identifier string
	Score      int
}

func Run(repo repository.Repository, rules []Rule) []Result {
	for _, detector := range collection {
		detector.Init(repo)
	}

	resultChannel := make(chan Result)
	for _, rule := range rules {
		for _, detector := range collection {
			if detector.Supports(rule) {
				go detector.Detect(repo, resultChannel, rule)
			}
		}
	}

	var results []Result
	for i := 0; i < len(rules); i++ {
		results = append(results, <-resultChannel)
	}
	return results
}

func register(detector Detector) {
	collection = append(collection, detector)
}
