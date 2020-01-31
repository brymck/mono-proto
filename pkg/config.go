package pkg

type Config struct {
	Prototool string       `yaml:"prototool,omitempty"`
	Repos     []RepoConfig `yaml:"repos,omitempty"`
	Language  string       `yaml:"language,omitempty"`
}

type RepoConfig struct {
	GitHub Repo `yaml:"github,omitempty"`
}

type Repo struct {
	Owner string `yaml:"owner,omitempty"`
	Name  string `yaml:"name,omitempty"`
}
