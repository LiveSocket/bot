package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	projects, err := getProjects(cwd)
	if err != nil {
		panic(err)
	}

	for _, project := range projects {
		err := dockerBuild(project)
		if err != nil {
			panic(err)
		}
	}
}

func getProjects(dir string) ([]string, error) {
	if len(os.Args) > 1 && os.Args[1] != "" {
		return []string{os.Args[1]}, nil
	}

	var result []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".docker" {
			result = append(result, strings.Split(filepath.Base(path), ".")[0])
		}
		return nil
	})
	return result, err
}

func dockerBuild(project string) error {
	cmd := exec.Command("docker", "build", "--target", "release", "-t", "livesocket/"+project, "-f", project+".docker", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
