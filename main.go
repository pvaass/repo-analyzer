package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pvaass/repo-analyzer/pkg/analyze"
	"github.com/pvaass/repo-analyzer/pkg/discovery/detectors"
	"github.com/pvaass/repo-analyzer/pkg/repository"
	"github.com/pvaass/repo-analyzer/pkg/repository/platforms"
	"github.com/spf13/viper"
)

func main() {
	log.SetPrefix("[repo-analyzer] ")
	initConfig()

	if len(os.Args) < 2 {
		panic(fmt.Errorf("missing arg 'repo'"))
	}

	platformCollection := getPlatforms()

	platform, err := platformCollection.ForURI(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	repository := repository.New(platform, os.Args[1])

	rules := detectors.GetRules(viper.GetString("rules.path"))
	if len(rules) < 1 {
		fmt.Println("No rules found in configured directory " + viper.GetString("rules.path") + ", exiting.")
		os.Exit(1)
	}

	results := analyze.Run(repository, rules)
	for _, result := range results {
		fmt.Println(fmt.Sprintf("%s: %d", result.Identifier, result.Score))
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/repo-analyzer")
	viper.AddConfigPath("$HOME/.repo-analyzer")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error while reading the config file: %s", err))
	}

	viper.SetDefault("github.enable", true)
	viper.SetDefault("bitbucket.enable", true)
	viper.SetDefault("rules.path", "/etc/repo-analyzer")

	if viper.GetBool("github.enable") && !viper.IsSet("github.token") {
		log.Println("GitHub is enabled, but no token was configured. Will not be able to access private repositories")
	}
}

func getPlatforms() platforms.Platforms {
	var enabledPlatforms []platforms.Platform

	if viper.GetBool("github.enable") {
		enabledPlatforms = append(enabledPlatforms, &platforms.GitHub{
			Token: viper.GetString("github.token"),
		})
	}
	if viper.GetBool("bitbucket.enable") {
		enabledPlatforms = append(enabledPlatforms, &platforms.BitBucket{})
	}

	if len(enabledPlatforms) < 1 {
		fmt.Println("At least one platform needs to be enabled. Exiting.")
		os.Exit(1)
	}

	collection := platforms.Platforms{}
	collection.Add(enabledPlatforms...)
	return collection
}
