package git

import (
	"github.com/elissonalvesilva/cortex/internal/cobrax"
)

var (
	varBranch string

	varCloneDir  string
	varCloneRepo string

	varPullWithRebase bool

	varCreateWithCheckout bool

	varCommitMessage string

	CloneCmd    = cobrax.NewCommand("clone", cobrax.WithRunE(cloneCommand))
	PullCmd     = cobrax.NewCommand("pull", cobrax.WithRunE(pullCommand))
	PushCmd     = cobrax.NewCommand("push", cobrax.WithRunE(pushCommand))
	CheckoutCmd = cobrax.NewCommand("checkout", cobrax.WithRunE(checkoutCommand))
	CommitCmd   = cobrax.NewCommand("commit", cobrax.WithRunE(commitCommand))
	BranchCmd   = cobrax.NewCommand("branch", cobrax.WithRunE(branchCommand))
	StatusCmd   = cobrax.NewCommand("status", cobrax.WithRunE(statusCommand))
	AddCmd      = cobrax.NewCommand("add", cobrax.WithRunE(addCommand))
)

func init() {
	CloneCmd.Flags().StringVarPWithDefaultValue(&varCloneDir, "dir", "d", "")
	CloneCmd.Flags().StringVarPWithDefaultValue(&varCloneRepo, "repo", "r", "")

	PullCmd.Flags().BoolVarPWithDefaultValue(&varPullWithRebase, "rebase", "r", false)

	PushCmd.Flags().StringVarPWithDefaultValue(&varBranch, "branch", "b", "")

	CheckoutCmd.Flags().BoolVarPWithDefaultValue(&varCreateWithCheckout, "create", "c", false)
	CheckoutCmd.Flags().StringVarPWithDefaultValue(&varBranch, "branch", "b", "")

	CommitCmd.Flags().StringVarPWithDefaultValue(&varCommitMessage, "message", "m", "")

	BranchCmd.Flags().StringVarPWithDefaultValue(&varBranch, "branch", "b", "")
}
