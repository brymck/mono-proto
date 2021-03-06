package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/brymck/mono-proto/internal"
	"github.com/brymck/mono-proto/pkg"
)

func readConfig(configPath string) (*pkg.Config, error) {
	var cfg pkg.Config
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &cfg)
	return &cfg, err
}

func main() {
	logger := log.New(os.Stdout, "", 0)
	env, err := internal.NewEnvironment(logger)
	if err != nil {
		log.Fatal(err)
	}

	configPath := flag.String("config", ".mono-proto.yaml", "path to mono-proto.yaml configuration file")
	env.RootDirectory = *flag.String("schemas-root", env.RootDirectory, "path to schemas root directory")
	flag.Parse()

	cfg, err := readConfig(*configPath)
	if err != nil {
		logger.Fatal(err)
	}
	ghs := make([]*internal.GitHub, len(cfg.Repos))
	for i, repoCfg := range cfg.Repos {
		gh := internal.NewGitHub(&repoCfg.GitHub, env, logger)
		ghs[i] = gh
	}
	for _, gh := range ghs {
		dir := env.GetLocalDirectory(gh)
		if err = gh.Sync(dir); err != nil {
			logger.Fatal(err)
		}
	}
	for _, gh := range ghs {
		dir := env.GetLocalDirectory(gh)
		if err = env.RunCommand(
			"docker",
			"run",
			"-v",
			fmt.Sprintf("%s:/work", dir),
			"brymck/prototool-java-typescript:1.8.1",
			"prototool",
			"all",
			"proto",
		); err != nil {
			logger.Fatal(err)
		}
	}
}
