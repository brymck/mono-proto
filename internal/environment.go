package internal

import (
	"log"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
)

type Environment struct {
	RootDirectory string
	logger        *log.Logger
}

type repository interface {
	Url() *url.URL
}

func NewEnvironment(logger *log.Logger) (*Environment, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	root := path.Join(home, ".mono-proto")
	return &Environment{RootDirectory: root, logger: logger}, nil
}

func (env *Environment) DirectoryExists(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return true, nil
	}
}

func (env *Environment) MakeDirectory(dir string) error {
	return os.MkdirAll(dir, 0755)
}

func (env *Environment) GetLocalDirectory(repo repository) string {
	u := repo.Url()
	p := strings.TrimSuffix(u.Path, ".git")
	return path.Join(env.RootDirectory, u.Host, p)
}

func (env *Environment) RunCommand(name string, arg ...string) error {
	env.logger.Println("running command:")
	env.logger.Printf("  %s %s\n", name, strings.Join(arg, " "))
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
