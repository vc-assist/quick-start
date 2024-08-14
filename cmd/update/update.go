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

func pullRepo(path, url string) {
	fmt.Printf("Pulling %s...\n", url)

	originalPath, err := os.Getwd()
	if err != nil {
		fatal(err)
	}
	err = os.Chdir(path)
	if err != nil {
		fatal(err)
	}

	out, err := exec.Command("git", "pull").
		CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		fatal(err)
	}
	out, err = exec.Command("git", "submodule", "update", "--init", "--recursive").
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
	pullRepo("./backend", "https://github.com/vc-assist/backend")
	pullRepo("./frontend", "https://github.com/vc-assist/frontend")
	pullRepo("./ui", "https://github.com/vc-assist/ui")
	pullRepo("./mobile", "https://github.com/vc-assist/mobile")
	pullRepo("./desktop", "https://github.com/vc-assist/desktop")
	pullRepo("./deployment", "https://github.com/vc-assist/deployment")

	installNodeModules("./frontend")
	installNodeModules("./desktop")
	installNodeModules("./desktop/frontend")
	installNodeModules("./ui")
}
