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
	searchEngineURL = "https://www.google.com/search?q=%s"
	videoURL        = "https://www.youtube.com/results?search_query=%s"
)

func searchCommand(_ *cobra.Command, _ []string) (err error) {
	defer func() {
		if err == nil {
			fmt.Println(color.Green.Render("search completed successfully\n"))
		}
	}()

	if err = openBrowser(defineSearch()); err != nil {
		return fmt.Errorf(color.Red.Render("error opening browser: %w", err))
	}

	return nil
}

func defineSearch() string {
	if varVideo {
		return fmt.Sprintf(buildSearchURL(videoURL, varQuery))
	}

	return fmt.Sprintf(buildSearchURL(searchEngineURL, varQuery))
}

func buildSearchURL(url, query string) string {
	return fmt.Sprintf(url, query)
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
