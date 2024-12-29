package search

import "github.com/elissonalvesilva/cortex/internal/cobrax"

var (
	varQuery string
	varVideo bool

	Cmd = cobrax.NewCommand("search", cobrax.WithRunE(searchCommand))
)

func init() {
	searchFlags := Cmd.Flags()
	searchFlags.StringVarP(&varQuery, "query", "q")
	searchFlags.BoolVarPWithDefaultValue(&varVideo, "video", "v", false)
}
