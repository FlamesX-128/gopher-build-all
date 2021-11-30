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

package handlers

import (
	"fmt"
	"os/exec"
	"path"
	"sync"

	"github.com/FlamesX-128/gopher-build-all/src/helpers"
)

func createBin(directory string, file string, flags Flags, system System) {
	baseName := fmt.Sprintf("%s-%s-%s", flags.Bin_name, system.Name, system.Arch)

	if flags.Sub_folder {
		baseName = path.Join(baseName, baseName)

	}

	err := exec.Command(
		"env",
		("GOOS=" + system.Name),
		("GOARCH=" + system.Arch),
		"go",
		"build",
		"-o",
		path.Join(directory, baseName),
		file,
	).Run()

	if err != nil {
		fmt.Printf("Failed to create file: %s-%s-%s\n", flags.Bin_name, system.Name, system.Arch)

	}
}

func BinaryHandler(baseDir string, file string, flags Flags, systems []System) {
	var (
		i  uint8
		wg sync.WaitGroup
	)

	for _, system := range systems {
		// Check if it should only be created on specific platforms.
		if flags.Only_systems != nil && !helpers.Contains(flags.Only_systems, system.Name) {

			continue
		}

		// Check the maximum amount of goroutines.
		if i == flags.Max_goruntines {
			wg.Wait()

			i = 0
		}

		wg.Add(1)
		i++

		go func(system System) {
			defer wg.Done()

			createBin(baseDir, file, flags, system)
		}(system)

	}
}
