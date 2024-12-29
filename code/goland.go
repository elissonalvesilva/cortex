package code

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
)

func golandCommand(_ *cobra.Command, _ []string) (err error) {
	defer func() {
		if err == nil {
			fmt.Println(color.Green.Render("GoLand opened successfully"))
		}
	}()

	if varDirectory != "" {
		return openDirectoryInGoLand(varDirectory)
	} else if varFile != "" {
		return openFileInGoLand(varFile)
	}

	if varDirectory == "" {
		return openGoLand()
	}

	return nil
}

func openGoLand() error {
	switch runtime.GOOS {
	case "windows":
		if err := exec.Command("goland").Run(); err != nil {
			return exec.Command("C:\\Program Files\\JetBrains\\GoLand\\bin\\goland64.exe").Start()
		}
	case "darwin":
		if err := exec.Command("goland").Run(); err != nil {
			return exec.Command("open", "-a", "GoLand").Start()
		}
	case "linux":
		if err := exec.Command("goland").Run(); err != nil {
			if _, err := exec.LookPath("/opt/GoLand/bin/goland"); err == nil {
				return exec.Command("/opt/GoLand/bin/goland").Start()
			}

			if _, err := exec.LookPath("/usr/local/bin/goland"); err == nil {
				return exec.Command("/usr/local/bin/goland").Start()
			}
		}

		return fmt.Errorf(color.Red.Render("not supported operating system"))
	}
	return nil
}

func openDirectoryInGoLand(directory string) error {
	if directory == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf(color.Red.Render("error getting current directory: %w", err))
		}
		directory = currentDir
	}

	switch runtime.GOOS {
	case "windows":
		if err := exec.Command("goland", directory).Run(); err != nil {
			return exec.Command("C:\\Program Files\\JetBrains\\GoLand\\bin\\goland64.exe", directory).Start()
		}
	case "darwin":
		if err := exec.Command("goland", directory).Run(); err != nil {
			return exec.Command("open", "-a", "GoLand", directory).Start()
		}
	case "linux":
		if err := exec.Command("goland", directory).Run(); err != nil {
			if _, err := exec.LookPath("/opt/GoLand/bin/goland"); err == nil {
				return exec.Command("/opt/GoLand/bin/goland", directory).Start()
			}

			if _, err := exec.LookPath("/usr/local/bin/goland"); err == nil {
				return exec.Command("/usr/local/bin/goland", directory).Start()
			}
		}

		return fmt.Errorf(color.Red.Render("not supported operating system"))
	}
	return nil
}

func openFileInGoLand(filePath string) error {
	if filePath == "" {
		return fmt.Errorf(color.Red.Render("no file path provided"))
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf(color.Red.Render("file does not exist: %s", filePath))
	}

	switch runtime.GOOS {
	case "windows":
		if err := exec.Command("goland", filePath).Run(); err != nil {
			return exec.Command("C:\\Program Files\\JetBrains\\GoLand\\bin\\goland64.exe", filePath).Start()
		}
	case "darwin":
		if err := exec.Command("goland", filePath).Run(); err != nil {
			return exec.Command("open", "-a", "GoLand", filePath).Start()
		}
	case "linux":
		if err := exec.Command("goland", filePath).Run(); err != nil {
			if _, err := exec.LookPath("/opt/GoLand/bin/goland"); err == nil {
				return exec.Command("/opt/GoLand/bin/goland", filePath).Start()
			}

			if _, err := exec.LookPath("/usr/local/bin/goland"); err == nil {
				return exec.Command("/usr/local/bin/goland", filePath).Start()
			}
		}

		return fmt.Errorf(color.Red.Render("GoLand not found or cannot open the file"))
	}

	return fmt.Errorf("not supported operating system")
}
