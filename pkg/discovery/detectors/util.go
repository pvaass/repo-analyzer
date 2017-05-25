package detectors

import (
	"github.com/pvaass/repo-analyzer/pkg/repository"
)

func findFile(repo repository.Repository, paths []string) []byte {
	for _, path := range paths {
		file, _ := repo.File(path)
		if len(file) > 0 {
			return file
		}
	}
	return []byte{}
}
