package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pvaass/repo-analyzer/pkg/analyze"
	"github.com/pvaass/repo-analyzer/pkg/repository"
	"github.com/pvaass/repo-analyzer/pkg/repository/platforms"
	"github.com/spf13/viper"
)

func main() {
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

	analyze.Run(repository)
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

	if viper.GetBool("github.enable") && !viper.IsSet("github.token") {
		log.Fatal("GitHub is enabled, but no token was configured")
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
