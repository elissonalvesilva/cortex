package code

import (
	"github.com/elissonalvesilva/cortex/internal/cobrax"
	"github.com/spf13/cobra"
)

var (
	varCodeType  string
	varDirectory string
	varFile      string

	Cmd        = cobrax.NewCommand("code", cobrax.WithRunE(codeCommand))
	vscodeCmd  = cobrax.NewCommand("vscode", cobrax.WithRunE(vscodeCommand))
	sublimeCmd = cobrax.NewCommand("sublime", cobrax.WithRunE(sublimeCommand))
	golandCmd  = cobrax.NewCommand("goland", cobrax.WithRunE(golandCommand))
)

func init() {
	Cmd.Flags().StringVarPWithDefaultValue(&varDirectory, "directory", "d", "")
	Cmd.Flags().StringVarPWithDefaultValue(&varFile, "file", "f", "")
	Cmd.Flags().StringVarPWithDefaultValue(&varCodeType, "type", "t", "vscode")

	Cmd.AddCommand(vscodeCmd, sublimeCmd, golandCmd)
}

func codeCommand(_ *cobra.Command, _ []string) error {
	switch varCodeType {
	case "vscode":
		return vscodeCommand(nil, nil)
	case "sublime":
		return sublimeCommand(nil, nil)
	case "goland":
		return golandCommand(nil, nil)
	default:
		return vscodeCommand(nil, nil)
	}
}
