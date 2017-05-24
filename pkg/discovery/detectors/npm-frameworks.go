package detectors

import "github.com/pvaass/repo-analyzer/pkg/repository"

type NodeFrameworkDetector struct {
	rule NodeFrameworkDetectionRule
}

type NodeFrameworkDetectionRule struct {
	Identifier  string
	PackageName string
}

func (f NodeFrameworkDetector) Detect(repo repository.Repository, resultChannel chan Result) {
	result := Result{Identifier: f.rule.Identifier}
	if !hasNpm(repo) {
		resultChannel <- result
		return
	}

	if npmRequiresPackage(repo.File("package.json"), f.rule.PackageName) {
		result.Score = 100
	}

	resultChannel <- result
}

func init() {
	create := func(a string, b string) *NodeFrameworkDetector {
		return &NodeFrameworkDetector{NodeFrameworkDetectionRule{a, b}}
	}

	register(create("js-react", "react"))
	register(create("js-express", "express"))
	register(create("js-angularjs", "angular"))
	register(create("js-vue", "vue"))

	// TODO:
	// ember.js
	// meteor.js
}
