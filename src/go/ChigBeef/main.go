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
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal("generate a data file using 'generate.py' before running this")
	}
	file := strings.Split(string(f), "\r")

	data := [][]uint16{}
	for i, l := range file {
		// Get numbers in string form (["2", "7", ...])
		line := strings.Split(l, ",")

		data = append(data, []uint16{})
		for _, n := range line {
			num, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println(err)
				fmt.Println(n)
				log.Fatal("issue with converting a string to a number, maybe check 'data.txt'?")
			}
			data[i] = append(data[i], uint16(num))
		}
	}
	// Data should be nice and full right now

	var minVal float64 = 1 << 32 // Some large number
	var maxVal float64 = 0       // Some small number
	var meanVal float64 = 0
	var totalVal float64 = 0
	var count float64 = 0

	for _, l := range data {
		var rawSum uint32 = 0
		for _, n := range l {
			rawSum += uint32(n)
		}
		sum := uint16(rawSum)
		// Now we have the sum of the line, we can start doing our extra caluclations.

		// y = x^(3/2)
		val := math.Pow(float64(sum), 3)
		val = math.Sqrt(val)
		val = math.Floor(val)

		// This is for the mean
		totalVal += val
		count++

		// Find new min and max
		if val < minVal {
			minVal = val
		}
		if val > maxVal {
			maxVal = val
		}
	}

	meanVal = math.Floor(totalVal / count)

	fmt.Println(
		"MIN",
		strconv.FormatFloat(minVal, 'f', 0, 64),
		"MAX",
		strconv.FormatFloat(maxVal, 'f', 0, 64),
		"MEAN",
		strconv.FormatFloat(meanVal, 'f', 0, 64),
	)
}
