package code

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
)

func vscodeCommand(_ *cobra.Command, _ []string) (err error) {
	defer func() {
		if err == nil {
			fmt.Println(color.Green.Render("vscode opened successfully"))
		}
	}()

	if varDirectory != "" {
		return openDirectoryInVSCode(varDirectory)
	} else if varDirectory == "" {
		return openVSCode()
	}

	if varFile != "" {
		return openFileInVSCode(varFile)
	}

	return nil
}

func openVSCode() error {
	switch runtime.GOOS {
	case "windows":
		if err := exec.Command("code").Run(); err != nil {
			return exec.Command("C:\\Program Files\\Microsoft VS Code\\Code.exe").Start()
		}
	case "darwin":
		if err := exec.Command("code").Run(); err != nil {
			return exec.Command("open", "-a", "Visual Studio Code").Start()
		}
	case "linux":
		if err := exec.Command("code").Run(); err != nil {
			if _, err := exec.LookPath("/snap/bin/code"); err == nil {
				return exec.Command("/snap/bin/code").Start()
			}

			if _, err := exec.LookPath("/usr/bin/code"); err == nil {
				return exec.Command("/usr/bin/code").Start()
			}

			if _, err := exec.LookPath("/opt/visual-studio-code/bin/code"); err == nil {
				return exec.Command("/opt/visual-studio-code/bin/code").Start()
			}
		}

		return fmt.Errorf(color.Red.Render("not supported operating system"))
	}
	return nil
}

func openDirectoryInVSCode(directory string) error {
	if directory == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf(color.Red.Render("error getting current directory: %w", err))
		}
		directory = currentDir
	}

	switch runtime.GOOS {
	case "windows":
		if err := exec.Command("code", directory).Run(); err != nil {
			return exec.Command("C:\\Program Files\\Microsoft VS Code\\Code.exe", directory).Start()
		}
	case "darwin":
		if err := exec.Command("code", directory).Run(); err != nil {
			return exec.Command("open", "-a", "Visual Studio Code", directory).Start()
		}
	case "linux":
		if err := exec.Command("code", directory).Run(); err != nil {
			if _, err := exec.LookPath("/snap/bin/code"); err == nil {
				return exec.Command("/snap/bin/code", directory).Start()
			}

			if _, err := exec.LookPath("/usr/bin/code"); err == nil {
				return exec.Command("/usr/bin/code", directory).Start()
			}

			if _, err := exec.LookPath("/opt/visual-studio-code/bin/code"); err == nil {
				return exec.Command("/opt/visual-studio-code/bin/code", directory).Start()
			}
		}

		return fmt.Errorf(color.Red.Render("not supported operating system"))
	}
	return nil
}

func openFileInVSCode(filePath string) error {
	if filePath == "" {
		return fmt.Errorf(color.Red.Render("no file path provided"))
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf(color.Red.Render("file does not exist: %s", filePath))
	}

	switch runtime.GOOS {
	case "windows":
		if err := exec.Command("code", filePath).Run(); err != nil {
			return exec.Command("C:\\Program Files\\Microsoft VS Code\\Code.exe", filePath).Start()
		}
	case "darwin":
		if err := exec.Command("code", filePath).Run(); err != nil {
			return exec.Command("open", "-a", "Visual Studio Code", filePath).Start()
		}
	case "linux":
		if err := exec.Command("code", filePath).Run(); err != nil {
			if _, err := exec.LookPath("/snap/bin/code"); err == nil {
				return exec.Command("/snap/bin/code", filePath).Start()
			}

			if _, err := exec.LookPath("/usr/bin/code"); err == nil {
				return exec.Command("/usr/bin/code", filePath).Start()
			}

			if _, err := exec.LookPath("/opt/visual-studio-code/bin/code"); err == nil {
				return exec.Command("/opt/visual-studio-code/bin/code", filePath).Start()
			}
		}
		return fmt.Errorf(color.Red.Render("VSCode not found or cannot open the file"))
	}

	return fmt.Errorf("not supported operating system")
}
