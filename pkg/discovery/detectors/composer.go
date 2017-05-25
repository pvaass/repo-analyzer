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
	return len(getComposer(repo)) > 0
}

func getComposer(repo repository.Repository) []byte {
	return findFile(
		repo,
		[]string{
			"composer.json",
			"app/composer.json",
		},
	)
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
