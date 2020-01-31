package pkg

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"gopkg.in/yaml.v2"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    Config
	}{
		{
			name:    "Prototool",
			content: "prototool: foo",
			want:    Config{Prototool: "foo"},
		},
		{
			name:    "Repo",
			content: "github:\n  owner: brymck\n  name: mono-proto",
			want:    Config{GitHub: Repo{Owner: "brymck", Name: "mono-proto"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Config{}
			err := yaml.Unmarshal([]byte(tt.content), &cfg)
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(cfg, tt.want) {
				t.Errorf("want %+v; got %+v", tt.want, cfg)
			}

		})
	}
}
