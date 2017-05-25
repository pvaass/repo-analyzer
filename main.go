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

	platforms := getPlatforms()
	platform := platforms.ForURI(os.Args[1])
	repository := repository.New(platform, os.Args[1])
	analyze.Run(repository)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error while reading the config file: %s", err))
	}

	if !viper.IsSet("github.token") {
		log.Println("No github token set. Will not use github as repository source.")
	}
}

func getPlatforms() platforms.Platforms {
	collection := platforms.Platforms{}
	github := &platforms.GitHub{
		Token: viper.GetString("github.token"),
	}

	bitbucket := &platforms.BitBucket{}
	collection.Add(github)
	collection.Add(bitbucket)
	return collection
}
