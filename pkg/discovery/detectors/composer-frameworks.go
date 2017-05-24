package detectors

import "github.com/pvaass/repo-analyzer/pkg/repository"

type ComposerFrameworkDetector struct {
	rule ComposerFrameworkDetectionRule
}

type ComposerFrameworkDetectionRule struct {
	Identifier  string
	PackageName string
}

func (f ComposerFrameworkDetector) Detect(repo repository.Repository, resultChannel chan Result) {
	result := Result{Identifier: f.rule.Identifier}
	if !hasComposer(repo) {
		resultChannel <- result
		return
	}

	if composerRequiresPackage(repo.File("composer.json"), f.rule.PackageName) {
		result.Score = 100
	}

	resultChannel <- result
}

func init() {
	create := func(a string, b string) *ComposerFrameworkDetector {
		return &ComposerFrameworkDetector{ComposerFrameworkDetectionRule{a, b}}
	}

	register(create("php-symfony", "symfony/symfony"))
	register(create("php-laravel", "laravel/framework"))
}
