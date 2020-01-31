package mocks

type Environment struct{
	Spy [][]string
}

func (env *Environment) DirectoryExists(dir string) (bool, error) {
	env.Spy = append(env.Spy, []string{"DirectoryExists", dir})
	if dir == "exists" {
		return true, nil
	} else {
		return false, nil
	}
}

func (env *Environment) MakeDirectory(dir string) error {
	env.Spy = append(env.Spy, []string{"MakeDirectory", dir})
	return nil
}

func (env *Environment) RunCommand(name string, arg ...string) error {
	call := []string{"RunCommand", name}
	call = append(call, arg...)
	env.Spy = append(env.Spy, call)
	return nil
}
