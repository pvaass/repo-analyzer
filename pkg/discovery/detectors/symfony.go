package detectors

import "github.com/pvaass/repo-analyzer/pkg/repository"

type Symfony struct{}

func (f Symfony) Detect(repo repository.Repository, resultChannel chan Result) {
	result := Result{
		Identifier: "symfony",
	}
	if !hasComposer(repo) {
		resultChannel <- result
		return
	}

	if composerRequiresPackage(repo.File("composer.json"), "symfony/symfony") {
		result.Score = 100
	}

	resultChannel <- result
}

func init() {
	register(Symfony{})
}
