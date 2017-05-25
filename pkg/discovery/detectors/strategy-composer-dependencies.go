package detectors

import (
	"encoding/json"
	"log"

	"github.com/pvaass/repo-analyzer/pkg/repository"
)

type ComposerDependencyDetector struct {
	composer ComposerDependencyMap
}

type ComposerDependencyMap struct {
	Require map[string]string
}

func (f ComposerDependencyDetector) Supports(rule Rule) bool {
	return rule.Strategy == "composer#d"
}

func (f *ComposerDependencyDetector) Init(repo repository.Repository) {
	file := findFile(
		repo,
		[]string{
			"composer.json",
			"app/composer.json",
		},
	)
	if len(file) <= 0 {
		return
	}

	err := json.Unmarshal(file, &f.composer)
	if err != nil {
		log.Panic("Invalid Json Decode", err)
	}
}

func (f ComposerDependencyDetector) Detect(repo repository.Repository, resultChannel chan Result, rule Rule) {
	result := Result{Identifier: rule.Name}

	_, ok := f.composer.Require[rule.Arguments[0]]
	if ok {
		result.Score = 100
	}

	resultChannel <- result
}

func init() {
	register(&ComposerDependencyDetector{})
}
