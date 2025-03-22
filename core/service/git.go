package service

import (
	"errors"
	"os/exec"
)

type Git struct {
}

func NewGit() (*Git, error) {
	checkCmd := exec.Command(gitRootCommand, "rev-parse", "--is-inside-work-tree")
	if err := checkCmd.Run(); err != nil {
		return nil, errors.New("not a git repository")
	}
	return &Git{}, nil
}

var gitRootCommand = "git"

func (g *Git) Diff() ([]byte, error) {

	arg1 := "diff"
	arg2 := "--cached"
	arg3 := "--minimal"
	cmd := exec.Command(gitRootCommand, arg1, arg2, arg3)

	stdout, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return stdout, nil
}
