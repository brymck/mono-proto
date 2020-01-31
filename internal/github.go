package internal

import (
	"fmt"
	"log"
	"net/url"
	"path"
	"path/filepath"

	"github.com/brymck/mono-proto/pkg"
)

type environment interface {
	DirectoryExists(string) (bool, error)
	MakeDirectory(string) error
	RunCommand(string, ...string) error
}

type GitHub struct {
	Username string
	Repository string
	environment environment
	logger *log.Logger
}

func NewGitHub(repo *pkg.Repo, env environment, logger *log.Logger) *GitHub {
	return &GitHub{
		Username: repo.Owner,
		Repository: repo.Name,
		environment: env,
		logger: logger,
	}
}

func (gh *GitHub) shallowClone(dir string) error {
	return gh.environment.RunCommand("git", "clone", "--depth", "1", gh.Url().String(), dir)
}

func (gh *GitHub) Url() *url.URL {
	p := fmt.Sprintf("%s/%s.git", gh.Username, gh.Repository)
	return &url.URL{
		Scheme: "https",
		Host: "github.com",
		Path: p,
	}
}

func (gh *GitHub) pull(dir string) error {
	dotGit := path.Join(dir, ".git")
	if err := gh.environment.RunCommand("git", "--work-tree", dir, "--git-dir", dotGit, "pull"); err != nil {
		return fmt.Errorf("error updating repo in %s: %v", dir, err)
	}
	return nil
}

func (gh *GitHub) Sync(dir string) error {
	gh.logger.Printf("cloning from %s to %s", gh.Url(), dir)
	exists, err := gh.environment.DirectoryExists(dir)
	if err != nil {
		return err
	}
	if exists {
		return gh.pull(dir)
	} else {
		parent := filepath.Dir(dir)
		if err = gh.environment.MakeDirectory(parent); err != nil {
			return err
		}
		return gh.shallowClone(dir)
	}
}
