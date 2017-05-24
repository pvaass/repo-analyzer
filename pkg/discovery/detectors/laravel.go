package detectors

import (
	"encoding/json"
	"log"

	"github.com/pvaass/repo-analyzer/pkg/repository"
)

type Laravel struct {
	composer struct {
		Require struct {
			Laravel string `json:"laravel/framework"`
		} `json:"require"`
	}
}

func (f Laravel) Detect(repo repository.Repository, resultChannel chan Result) {
	result := Result{
		Identifier: "laravel",
		Score:      0,
	}
	if !hasComposer(repo) {
		resultChannel <- result
		return
	}

	f.getComposer(repo)

	if f.composer.Require.Laravel != "" {
		result.Score = 100
	}

	resultChannel <- result
}

func (f *Laravel) getComposer(repo repository.Repository) {
	file := repo.File("composer.json")

	err := json.Unmarshal([]byte(file), &f.composer)
	if err != nil {
		log.Panic("Invalid Json Decode", err)
	}

}

func init() {
	register(Laravel{})
}
