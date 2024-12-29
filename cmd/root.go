package cmd

import (
	"fmt"
	"github.com/gookit/color"
	"os"
	"runtime"

	"github.com/elissonalvesilva/cortex/internal/cobrax"
	"github.com/elissonalvesilva/cortex/internal/version"
	"github.com/elissonalvesilva/cortex/search"
)

const (
	codeFailure = 1
)

var (
	rootCmd = cobrax.NewCommand("cortex")
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(color.Red.Render(err.Error()))
		os.Exit(codeFailure)
	}
}

func init() {
	rootCmd.Version = fmt.Sprintf(
		"%s %s/%s", version.BuildVersion,
		runtime.GOOS, runtime.GOARCH)
	rootCmd.AddCommand(search.Cmd)
}
