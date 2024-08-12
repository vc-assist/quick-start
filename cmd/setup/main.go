package main

import (
	"fmt"
	"os"
	"os/exec"
)

func fatal(msg any) {
	fmt.Println(msg)
	os.Exit(1)
}

func cloneRepo(url string) {
	fmt.Printf("Cloning %s...\n", url)
	out, err := exec.Command("git", "clone", "--recurse-submodules", url).
		CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		fatal(err)
	}
}

func installNodeModules(path string) {
	fmt.Printf("Installing node_modules (%s)...\n", path)

	originalPath, err := os.Getwd()
	if err != nil {
		fatal(err)
	}
	err = os.Chdir(path)
	if err != nil {
		fatal(err)
	}

	out, err := exec.Command("pnpm", "install").
		CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		fatal(err)
	}

	err = os.Chdir(originalPath)
	if err != nil {
		fatal(err)
	}
}

func main() {
	cloneRepo("https://github.com/vc-assist/backend")
	cloneRepo("https://github.com/vc-assist/frontend")
	cloneRepo("https://github.com/vc-assist/ui")
	cloneRepo("https://github.com/vc-assist/mobile")
	cloneRepo("https://github.com/vc-assist/desktop")
	cloneRepo("https://github.com/vc-assist/deployment")

	installNodeModules("./frontend")
	installNodeModules("./desktop")
	installNodeModules("./desktop/frontend")
	installNodeModules("./ui")
}
