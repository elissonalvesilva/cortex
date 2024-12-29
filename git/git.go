package git

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

func cloneCommand(_ *cobra.Command, _ []string) error {
	dir := varCloneDir
	if dir == "" {
		dir = "."
	}

	return runGitCommand("clone", varCloneRepo, dir)
}

func pullCommand(_ *cobra.Command, _ []string) error {
	branch := varBranch
	if branch == "" {
		defaultBranch, err := getDefaultBranch()
		if err != nil {
			return err
		}
		branch = defaultBranch
	}

	if varPullWithRebase {
		return runGitCommand("pull", "--rebase", "origin", branch)
	}

	return runGitCommand("pull", "origin", branch)
}

func pushCommand(_ *cobra.Command, _ []string) error {
	branch := varBranch
	if branch == "" {
		defaultBranch, err := getDefaultBranch()
		if err != nil {
			return err
		}
		branch = defaultBranch
	}

	return runGitCommand("push", "origin", branch)
}

func checkoutCommand(_ *cobra.Command, _ []string) error {
	branch := varBranch
	if branch == "" {
		defaultBranch, err := getDefaultBranch()
		if err != nil {
			return err
		}
		branch = defaultBranch
	}

	if varCreateWithCheckout {
		return runGitCommand("checkout", "-b", branch)
	}

	return runGitCommand("checkout", branch)
}

func commitCommand(_ *cobra.Command, _ []string) error {
	message := varCommitMessage
	if message == "" {
		return fmt.Errorf(color.Red.Render("message is required"))
	}

	return runGitCommand("commit", "-m", message)
}

func statusCommand(_ *cobra.Command, args []string) error {
	return runGitCommand(append([]string{"status"}, args...)...)
}

func branchCommand(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return runGitCommand("branch")
	}

	return runGitCommand(append([]string{"branch"}, args...)...)
}

func addCommand(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		args = []string{"."}
	}

	return runGitCommand(append([]string{"add"}, args...)...)
}

func getDefaultBranch() (string, error) {
	cmd := exec.Command("git", "remote", "show", "origin")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error fetching default branch: %w", err)
	}

	for _, line := range strings.Split(string(output), "\n") {
		if strings.Contains(line, "HEAD branch") {
			parts := strings.Fields(line)
			if len(parts) > 2 {
				return parts[2], nil
			}
		}
	}

	return "", fmt.Errorf(color.Red.Render("error fetching default branch"))
}

func runGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running git command: %v\nOutput: %s", err, string(out))
	}

	fmt.Println(string(out))
	return nil
}
