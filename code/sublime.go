package code

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
)

func sublimeCommand(_ *cobra.Command, _ []string) (err error) {
	defer func() {
		if err == nil {
			fmt.Println(color.Green.Render("Sublime Text opened successfully"))
		}
	}()

	if varDirectory != "" {
		return openDirectoryInSublime(varDirectory)
	} else if varFile != "" {
		return openFileInSublime(varFile)
	}

	if varDirectory == "" {
		return openSublime()
	}

	return nil
}

func openSublime() error {
	switch runtime.GOOS {
	case "windows":
		if err := exec.Command("subl").Run(); err != nil {
			return exec.Command("C:\\Program Files\\Sublime Text 3\\sublime_text.exe").Start()
		}
	case "darwin":
		if err := exec.Command("subl").Run(); err != nil {
			return exec.Command("open", "-a", "Sublime Text").Start()
		}
	case "linux":
		if err := exec.Command("subl").Run(); err != nil {
			if _, err := exec.LookPath("/opt/sublime_text/sublime_text"); err == nil {
				return exec.Command("/opt/sublime_text/sublime_text").Start()
			}

			if _, err := exec.LookPath("/usr/bin/subl"); err == nil {
				return exec.Command("/usr/bin/subl").Start()
			}
		}

		return fmt.Errorf(color.Red.Render("not supported operating system"))
	}
	return nil
}

func openDirectoryInSublime(directory string) error {
	if directory == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf(color.Red.Render("error getting current directory: %w", err))
		}
		directory = currentDir
	}

	switch runtime.GOOS {
	case "windows":
		if err := exec.Command("subl", directory).Run(); err != nil {
			return exec.Command("C:\\Program Files\\Sublime Text 3\\sublime_text.exe", directory).Start()
		}
	case "darwin":
		if err := exec.Command("subl", directory).Run(); err != nil {
			return exec.Command("open", "-a", "Sublime Text", directory).Start()
		}
	case "linux":
		if err := exec.Command("subl", directory).Run(); err != nil {
			if _, err := exec.LookPath("/opt/sublime_text/sublime_text"); err == nil {
				return exec.Command("/opt/sublime_text/sublime_text", directory).Start()
			}

			if _, err := exec.LookPath("/usr/bin/subl"); err == nil {
				return exec.Command("/usr/bin/subl", directory).Start()
			}
		}

		return fmt.Errorf(color.Red.Render("not supported operating system"))
	}
	return nil
}

func openFileInSublime(filePath string) error {
	if filePath == "" {
		return fmt.Errorf(color.Red.Render("no file path provided"))
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf(color.Red.Render("file does not exist: %s", filePath))
	}

	switch runtime.GOOS {
	case "windows":
		if err := exec.Command("subl", filePath).Run(); err != nil {
			return exec.Command("C:\\Program Files\\Sublime Text 3\\sublime_text.exe", filePath).Start()
		}
	case "darwin":
		if err := exec.Command("subl", filePath).Run(); err != nil {
			return exec.Command("open", "-a", "Sublime Text", filePath).Start()
		}
	case "linux":
		if err := exec.Command("subl", filePath).Run(); err != nil {
			if _, err := exec.LookPath("/opt/sublime_text/sublime_text"); err == nil {
				return exec.Command("/opt/sublime_text/sublime_text", filePath).Start()
			}

			if _, err := exec.LookPath("/usr/bin/subl"); err == nil {
				return exec.Command("/usr/bin/subl", filePath).Start()
			}
		}

		return fmt.Errorf(color.Red.Render("Sublime Text not found or cannot open the file"))
	}

	return fmt.Errorf("not supported operating system")
}
