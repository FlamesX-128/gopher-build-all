/*
This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <https://unlicense.org>
*/
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

func contains(list []string, subStr string) bool {
	for _, str := range list {
		if subStr == str {
			return true

		}
	}

	return false
}

func main() {
	projectDir := flag.String("project-directory", "", "Path to the project")
	binaryName := flag.String("binary-name", "main", "Binary name when building")
	goruntines := flag.Int("max-goruntines", 3, "Maximum quantity of goroutines.")
	onlyIn := flag.String("onlyIn", "", "Build only on these systems")
	flag.Parse()

	if projectDir == nil || *projectDir == "" {
		di, err := os.Getwd()

		if err != nil {
			log.Panicln(err)

		}

		projectDir = &di
	}

	var onlyInSystems []string

	if onlyIn != nil && *onlyIn != "" {
		onlyInSystems = strings.Split(*onlyIn, " ")

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
		if (onlyInSystems != nil) && (len(onlyInSystems) > 0) {
			if !contains(onlyInSystems, system.Name) {

				continue
			}

		}

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
