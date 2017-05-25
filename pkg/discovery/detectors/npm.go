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
	return len(getNpm(repo)) > 0
}

func getNpm(repo repository.Repository) []byte {
	for _, file := range repo.List("") {
		if file.Name == "package.json" {
			return repo.File("package.json")
		}
	}

	for _, file := range repo.List("app") {
		if file.Name == "package.json" {
			return repo.File("app/package.json")
		}
	}

	var nothing []byte
	return nothing
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
