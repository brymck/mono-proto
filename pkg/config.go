package pkg

type Config struct {
	Prototool string `yaml:"prototool,omitempty"`
	GitHub    Repo   `yaml:"github,omitempty"`
	Language  string `yaml:"language,omitempty"`
}

type Repo struct {
	Owner string `yaml:"owner,omitempty"`
	Name  string `yaml:"name,omitempty"`
}
