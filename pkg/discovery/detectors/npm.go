package detectors

import (
	"encoding/json"
	"log"

	"github.com/pvaass/repo-analyzer/pkg/repository"
)

type Npm struct{}

func (p Npm) Detect(repo repository.Repository, resultChannel chan Result) {
	result := Result{
		Identifier: "npm",
	}
	if hasNpm(repo) {
		result.Score = 100
	}
	resultChannel <- result
}

func hasNpm(repo repository.Repository) bool {
	for _, name := range repo.FileNames() {
		if name == "package.json" {
			return true
		}
	}

	return false
}

func npmRequiresPackage(file []byte, packageName string) bool {
	var npm struct {
		Dependencies map[string]string
	}
	err := json.Unmarshal(file, &npm)
	if err != nil {
		log.Panic("Invalid Json Decode", err)
	}
	_, ok := npm.Dependencies[packageName]
	return ok
}

func init() {
	register(Npm{})
}
