// Copyright 2024 The original author
//
//	Licensed under the Apache License, Version 2.0 (the "License");
//	you may not use this file except in compliance with the License.
//	You may obtain a copy of the License at
//
//	    http://www.apache.org/licenses/LICENSE-2.0
//
//	Unless required by applicable law or agreed to in writing, software
//	distributed under the License is distributed on an "AS IS" BASIS,
//	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	See the License for the specific language governing permissions and
//	limitations under the License.

package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Get the command
	args := os.Args
	if len(args) != 2 {
		log.Fatal("give me something to time!")
	}
	command := args[1]

	cmd := exec.Command(command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	start := time.Now()

	// Run the code
	err := cmd.Start()
	if err != nil {
		log.Fatal("couldn't execute")
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal("couldn't execute")
	}

	end := time.Now()

	// Show how long it took
	elapsed := end.UnixMilli() - start.UnixMilli()
	log.Println("executed in", elapsed, "ms")
}
