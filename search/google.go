package search

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"os/exec"
	"runtime"
)

type (
	SearchCommand struct {
		Query string
	}
)

const (
	searchURL = "https://www.google.com/search?q=%s"
)

func searchCommand(_ *cobra.Command, _ []string) (err error) {
	defer func() {
		if err == nil {
			fmt.Println(color.Green.Render("search completed successfully\n"))
		}
	}()
	
	if err = openBrowser(buildSearchURL(varQuery)); err != nil {
		return fmt.Errorf(color.Red.Render("error opening browser: %w", err))
	}

	return nil
}

func buildSearchURL(query string) string {
	return fmt.Sprintf(searchURL, query)
}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	case "linux":
		cmd = "xdg-open"
		args = []string{url}
	default:
		return fmt.Errorf("not supported operating system")
	}

	return exec.Command(cmd, args...).Start()
}
