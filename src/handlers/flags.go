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
	"flag"
	"log"
	"os"
	"strings"
)

func FlagsHandler() Flags {
	bin_folder_name := flag.String(
		"bin-folder-name", "bin", "The name of the folder where the binaries will be placed.",
	)

	bin_name := flag.String(
		"bin-name", "main", "Name given to the binary.",
	)

	max_goruntines := flag.Int(
		"max-goruntines", 3, "Used to indicate the maximum number of goroutines.",
	)

	only_systems := flag.String(
		"only-systems", "", "Used to indicate to only build on specific operating systems.",
	)

	project_directory := flag.String(
		"project-directory", "", "It is used to indicate which is the project path.",
	)

	sub_folder := flag.Bool(
		"sub-folder", true, "Create a folder where the program puts the binary and static resources",
	)

	flag.Parse()

	if *max_goruntines < 1 {
		log.Fatalln("The maximum number of goroutines must be greater than 0")

	}

	if project_directory == nil || *project_directory == "" {
		directory, err := os.Getwd()

		if err != nil {
			log.Fatalln("Could not get the current directory")

		}

		project_directory = &directory
	}

	var only_systems_slice []string

	if only_systems != nil && *only_systems != "" {
		only_systems_slice = strings.Split(*only_systems, " ")

	}

	return Flags{
		Bin_folder_name:   *bin_folder_name,
		Bin_name:          *bin_name,
		Max_goruntines:    uint8(*max_goruntines),
		Only_systems:      only_systems_slice,
		Project_directory: *project_directory,
		Sub_folder:        *sub_folder,
	}
}
