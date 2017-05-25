package analyze

import (
	"sort"

	"github.com/pvaass/repo-analyzer/pkg/discovery/detectors"
	"github.com/pvaass/repo-analyzer/pkg/repository"
)

type ByName []detectors.Result

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Identifier < a[j].Identifier }

// Run analyzes a Repository for known languages and frameworks
func Run(repo repository.Repository, rules []detectors.Rule) []detectors.Result {
	results := detectors.Run(repo, rules)
	sort.Sort(ByName(results))

	return results
}
