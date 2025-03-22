package service

import (
	"errors"
	"fmt"
	"os/exec"
)

type CommitType string

const (
	CommitTypeFeat  CommitType = "feat"
	CommitTypeFix   CommitType = "fix"
	CommitTypeChore CommitType = "chore"
	CommitTypeDocs  CommitType = "docs"
	CommitTypeTest  CommitType = "test"
)

func IsValidCommitType(commitType string) bool {
	switch CommitType(commitType) {
	case CommitTypeFeat, CommitTypeFix, CommitTypeChore, CommitTypeDocs, CommitTypeTest:
		return true
	default:
		return false
	}
}

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

func (g *Git) IsStaged() error {
	checkCmd := exec.Command(gitRootCommand, "diff", "--cached", "--quiet")
	if err := checkCmd.Run(); err == nil {
		return errors.New("no staged changes to commit")
	} else if exitError, ok := err.(*exec.ExitError); !ok || exitError.ExitCode() != 1 {
		return errors.New("failed to check staged changes")
	}
	return nil
}

func (g *Git) Commit(message string) error {
	cmd := exec.Command(gitRootCommand, "commit", "-m", message)
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("Successfully pushed xD")

	return nil
}
