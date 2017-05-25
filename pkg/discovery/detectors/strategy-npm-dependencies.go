package detectors

import (
	"encoding/json"
	"log"

	"github.com/pvaass/repo-analyzer/pkg/repository"
)

type NodeDependencyDetector struct {
	npm NodeDependencyMap
}

type NodeDependencyMap struct {
	Dependencies map[string]string
}

func (f NodeDependencyDetector) Supports(rule Rule) bool {
	return rule.Strategy == "npm#d"
}

func (f *NodeDependencyDetector) Init(repo repository.Repository) {
	file := findFile(
		repo,
		[]string{
			"package.json",
			"app/package.json",
		},
	)
	if len(file) <= 0 {
		return
	}

	err := json.Unmarshal(file, &f.npm)
	if err != nil {
		log.Panic("Invalid Json Decode", err)
	}
}

func (f NodeDependencyDetector) Detect(repo repository.Repository, resultChannel chan Result, rule Rule) {
	result := Result{Identifier: rule.Name}

	_, ok := f.npm.Dependencies[rule.Arguments[0]]
	if ok {
		result.Score = 100
	}

	resultChannel <- result
}

func init() {
	register(&NodeDependencyDetector{})
}
