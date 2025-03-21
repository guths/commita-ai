package internal

import "os/exec"

type Git struct {
}

func NewGit() *Git {
	return &Git{}
}

var gitRootCommand = "git"

func (g *Git) Diff() ([]byte, error) {
	arg1 := "diff"
	arg2 := "--unified=3"
	arg3 := "HEAD~"

	cmd := exec.Command(gitRootCommand, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	return stdout, nil
}
