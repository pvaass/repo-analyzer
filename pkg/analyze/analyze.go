package analyze

import (
	"log"

	"github.com/pvaass/repo-analyzer/pkg/discovery/detectors"
	"github.com/pvaass/repo-analyzer/pkg/repository"
)

// Run analyzes a Repository for known languages and frameworks
func Run(repo repository.Repository) {
	log.Println(detectors.Run(repo))
}
