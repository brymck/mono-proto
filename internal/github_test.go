package internal

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/brymck/mono-proto/internal/mocks"
	"github.com/brymck/mono-proto/pkg"
)

func TestSync(t *testing.T) {
	tests := []struct {
		name      string
		directory string
		wantEnv   [][]string
	}{
		{
			name:      "ExistingDirectory",
			directory: "exists",
			wantEnv: [][]string{
				{"DirectoryExists", "exists"},
				{"RunCommand", "git", "--work-tree", "exists", "--git-dir", "exists/.git", "pull"},
			},
		},
		{
			name:      "NonExistingDirectory",
			directory: "doesNotExist",
			wantEnv: [][]string{
				{"DirectoryExists", "doesNotExist"},
				{"MakeDirectory", "."},
				{"RunCommand", "git", "clone", "--depth", "1", "https://github.com/foo/bar.git", "doesNotExist"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &pkg.Repo{Owner: "foo", Name: "bar"}
			env := &mocks.Environment{}
			logger := log.New(ioutil.Discard, "", 0)
			gh := NewGitHub(repo, env, logger)
			err := gh.Sync(tt.directory)
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(env.Spy, tt.wantEnv) {
				t.Errorf("want environment calls to equal %q, got %q", tt.wantEnv, env.Spy)
			}
		})
	}
}
