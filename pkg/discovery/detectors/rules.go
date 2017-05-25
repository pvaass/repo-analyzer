package detectors

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type RuleSet struct {
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Name      string   `json:"name"`
	Arguments []string `json:"args"`
	Strategy  string   `json:"strg"`
}

func getRules() []Rule {
	var ruleSets []RuleSet
	filepath.Walk("./pkg/discovery/detectors/rules/", func(path string, _ os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".json") {
			return nil
		}

		file, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		var ruleSet RuleSet
		err = json.Unmarshal(file, &ruleSet)
		if err != nil {
			return err
		}

		ruleSets = append(ruleSets, ruleSet)

		return nil
	})

	var allRules []Rule
	for _, ruleSet := range ruleSets {
		allRules = append(allRules, ruleSet.Rules...)
	}

	log.Print(allRules)
	return allRules
}
