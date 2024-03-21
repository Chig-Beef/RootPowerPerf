// Copyright 2024 The original author
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Use a different character if they want
	args := os.Args

	var line_character string
	if len(args) > 1 {
		line_character = string(args[1][0])
		if args[1] == "nl" {
			line_character = "\n"
		}
	} else {
		line_character = "\r"
	}

	count_lines := 1_000_000
	count_nums := 1_000

	fmt.Println("generating")

	data := [][]string{}

	for i := range count_lines {
		data = append(data, []string{})
		for range count_nums {
			n := uint16(rand.Uint32())
			data[i] = append(data[i], strconv.Itoa(int(n)))
		}
	}

	fmt.Println("joining lines")

	output := []string{}
	for i := range len(data) {
		output = append(output, strings.Join(data[i], ","))
	}

	fmt.Println("saving")

	err := os.WriteFile("../runner/data.txt", []byte(strings.Join(output, line_character)), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done")
}
