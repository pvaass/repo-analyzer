package detectors

import (
	"encoding/json"
	"log"

	"github.com/pvaass/repo-analyzer/pkg/repository"
)

type Composer struct{}

func (d Composer) Detect(repo repository.Repository, resultChannel chan Result) {
	result := Result{
		Identifier: "composer",
	}
	if hasComposer(repo) {
		result.Score = 100
	}
	resultChannel <- result
}

func hasComposer(repo repository.Repository) bool {
	for _, name := range repo.FileNames() {
		if name == "composer.json" {
			return true
		}
	}

	return false
}

func composerRequiresPackage(file []byte, packageName string) bool {
	var composer struct {
		Require map[string]string
	}
	err := json.Unmarshal(file, &composer)
	if err != nil {
		log.Panic("Invalid Json Decode", err)
	}
	_, ok := composer.Require[packageName]
	return ok
}

func init() {
	register(Composer{})
}
