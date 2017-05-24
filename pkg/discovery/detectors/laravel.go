package detectors

import "github.com/pvaass/repo-analyzer/pkg/repository"

type Laravel struct{}

func (f Laravel) Detect(repo repository.Repository, resultChannel chan Result) {
	result := Result{
		Identifier: "laravel",
		Score:      0,
	}
	if !hasComposer(repo) {
		resultChannel <- result
		return
	}

	if composerRequiresPackage(repo.File("composer.json"), "laravel/framework") {
		result.Score = 100
	}

	resultChannel <- result
}

func init() {
	register(Laravel{})
}
