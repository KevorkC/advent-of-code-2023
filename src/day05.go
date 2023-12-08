package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type categoryMap struct {
	mapName             string
	destinationToSource map[int]int
}

// seed-to-soil map:
// 50 98 2
// 52 50 48

// The worker function that parses the map and creates a map
func parseAndCreateMapWorker(id int, jobs <-chan []string, results chan<- categoryMap) {
	for job := range jobs { // job is taken from the pool of jobs
		var sdm categoryMap
		sdm.mapName = job[0]
		sdm.destinationToSource = make(map[int]int)
		for _, line := range job[1:] {
			var source int
			var destination int
			var rangeLength int
			for i, subStrNumber := range strings.Fields(line) {
				if i == 0 {
					source, _ = strconv.Atoi(subStrNumber)
				} else if i == 1 {
					destination, _ = strconv.Atoi(subStrNumber)
				} else if i == 2 {
					rangeLength, _ = strconv.Atoi(subStrNumber)
				}
			}
			// Now populate the map
			for i := 0; i < rangeLength; i++ {
				sdm.destinationToSource[destination+i] = source + i
			}
		}
		results <- sdm

	}
}

func main() {
	file, err := os.Open("files/day05")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileLines []string

	for scanner.Scan() {
		var line string = scanner.Text()
		fileLines = append(fileLines, line)
		// fmt.Println("Appending line:", line)
	}
	// Print the file lines
	// for _, line := range fileLines {
	// 	fmt.Println(line)

	// Parse the seeds into a list of integers
	var seeds []int // Saving the seeds for later
	for _, seed := range strings.Fields(fileLines[0][7:]) {
		seed, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Println("Error converting seed substring", seed, "to number:", err)
		}
		seeds = append(seeds, seed)
	}
	// fmt.Println("seeds:", seeds)

	// Defining the job and result channels
	jobs := make(chan []string, 7)       // Each worker will get a list of strings to parse and create a map from
	results := make(chan categoryMap, 7) // Each worker will return a map

	// Preparing the jobs for the workers
	var mapSlice [][]string // The slice of strings that will be sent to the workers
	var startIndex int      // The start of the current map
	var endIndex int        // The end of the current map
	for i, line := range fileLines[1:] {
		// If the line contains the word "map", it is the start of a new map
		if strings.Contains(line, "map") {
			startIndex = i
		} else if line == "" {
			endIndex = i
			mapSlice = append(mapSlice, fileLines[startIndex:endIndex])
		}
	}
	// for _, mapDefinition := range mapSlice {
	// 	for _, line := range mapDefinition {
	// 		fmt.Println(line)
	// 	}
	// }

	// Send jobs to the workers.
	for _, mapDefinition := range mapSlice {
		jobs <- mapDefinition
	}
	close(jobs)

	// Use amount of cores when the work is CPU bound. Otherwise use thread amount if I/O-bound, when involving -
	// operations like network calls or disk access, where the CPU is often waiting and not continuously processing.
	var numWorkers int = runtime.NumCPU()
	var wg sync.WaitGroup
	fmt.Println("Using", numWorkers, "workers")
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			parseAndCreateMapWorker(w, jobs, results)
		}(w)
	}
	// /*
	//  *
	//  * Seperator
	//  *
	//  */

	// Wait for all the workers in the WaitGroup to finish.
	wg.Wait()
	close(results)

	// Collect one result
	// var seedToSoilMap categoryMap = <-results
	// fmt.Println("seed-to-soil map:", seedToSoilMap.destinationToSource[50])
	// // for destination, source := range seedToSoilMap.destinationToSource {
	// // 	fmt.Println(destination, source)
	// // }

	// Optionally, collect all the results.
	var maps []categoryMap
	for r := range results {
		maps = append(maps, r)
	}

	// Print all the maps' names
	for _, m := range maps {
		fmt.Println("Map names:", m.mapName)
	}
}

// type categoryMap struct {
// 	mapName             string
// 	destinationToSource map[int]int
// }

/* Input example:
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
*/
