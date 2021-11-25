package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
)

type System struct {
	Name string
	Type string
}

func build(dir string, target string, name string, system System) {
	err := exec.Command(
		"env",
		("GOOS=" + system.Name),
		("GOARCH=" + system.Type),
		"go",
		"build",
		"-o",
		path.Join(dir, "dist", (name+"-"+system.Name+"-"+system.Type)),
		target,
	).Run()

	if err != nil {
		log.Println("System error:", system.Name+"-"+system.Type)

	}
}

func main() {
	projectDir := flag.String("project-directory", "", "Path to the project")
	binaryName := flag.String("binary-name", "main", "Binary name when building")
	goruntines := flag.Int("max-goruntines", 3, "Maximum quantity of goroutines.")

	flag.Parse()

	if projectDir == nil || *projectDir == "" {
		di, err := os.Getwd()

		if err != nil {
			log.Panicln(err)

		}

		projectDir = &di
	}

	os.RemoveAll(path.Join(*projectDir, "dist"))
	err := os.Mkdir(path.Join(*projectDir, "dist"), 0755)

	if err != nil {
		log.Panicln(err)

	}

	out, err := exec.Command("go", "tool", "dist", "list").Output()

	if err != nil {
		log.Panicln(err)

	}

	var systems []System

	for _, arch := range strings.Split(string(out), "\n") {
		if arch == "" {

			continue
		}

		system := strings.Split(arch, "/")

		systems = append(systems, System{
			Name: system[0],
			Type: system[1],
		})
	}

	var mainFile string
	folders := []string{"src", "lib", ""}

	for _, folder := range folders {
		_, err := os.Stat(path.Join(*projectDir, folder, "main.go"))

		if err == nil {
			mainFile = path.Join(*projectDir, folder, "main.go")

		}
	}

	if mainFile == "" {
		log.Panicln("Could not find main.go")

	}

	var wg sync.WaitGroup
	var i int

	for _, system := range systems {
		if i == *goruntines {
			wg.Wait()
			i = 0
		}

		wg.Add(1)
		i++
		go func(system System) {
			defer wg.Done()
			build(*projectDir, mainFile, *binaryName, system)
		}(system)
	}

	wg.Wait()
}
