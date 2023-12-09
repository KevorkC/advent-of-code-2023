package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type categoryMap struct {
	ID                  int
	mapName             string
	sourceToDestination map[int]int
}

// The worker function that parses the map and creates a map
func parseAndCreateMapWorker(id int, jobs <-chan []string, results chan<- categoryMap) {
	for job := range jobs { // job is taken from the pool of jobs
		var sdm categoryMap
		// fmt.Println("Worker", id, "is parsing the map", job)
		sdm.mapName = job[0]

		switch sdm.mapName {
		case "seed-to-soil map:":
			sdm.ID = 1
		case "soil-to-fertilizer map:":
			sdm.ID = 2
		case "fertilizer-to-water map:":
			sdm.ID = 3
		case "water-to-light map:":
			sdm.ID = 4
		case "light-to-temperature map:":
			sdm.ID = 5
		case "temperature-to-humidity map:":
			sdm.ID = 6
		case "humidity-to-location map:":
			sdm.ID = 7
		default:
			fmt.Println("Error - Unknown map name:", sdm.mapName)
		}

		sdm.sourceToDestination = make(map[int]int)
		for _, line := range job[1:] {
			var source int
			var destination int
			var rangeLength int
			for i, subStrNumber := range strings.Fields(line) {
				if i == 0 {
					destination, _ = strconv.Atoi(subStrNumber)
				} else if i == 1 {
					source, _ = strconv.Atoi(subStrNumber)
				} else if i == 2 {
					rangeLength, _ = strconv.Atoi(subStrNumber)
				}
			}
			// Now populate the map
			for i := 0; i < rangeLength; i++ {
				sdm.sourceToDestination[source+i] = destination + i
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
	fileLines = fileLines[1:] // Remove the seeds line from the fileLines

	// Defining the job and result channels
	jobs := make(chan []string, 7)       // Each worker will get a list of strings to parse and create a map from
	results := make(chan categoryMap, 7) // Each worker will return a map

	// Preparing the jobs for the workers
	// var mapSlice [][]string // The slice of strings that will be sent to the workers
	var currentMapStart int = 0
	for i, line := range fileLines {
		if line == "" || i == len(fileLines)-1 {
			// Send the current map to jobs
			endIndex := i
			if line != "" {
				endIndex = i + 1
			}
			if endIndex > currentMapStart {
				jobs <- fileLines[currentMapStart:endIndex]
			}
			currentMapStart = i + 1 // Update the start for the next map
		}
	}

	close(jobs)

	// Use amount of cores when the work is CPU bound. Otherwise use thread amount if I/O-bound, when involving -
	// operations like network calls or disk access, where the CPU is often waiting and not continuously processing.
	var numWorkers int = runtime.NumCPU()
	var wg sync.WaitGroup
	// fmt.Println("Using", numWorkers, "workers")
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			parseAndCreateMapWorker(w, jobs, results)
		}(w)
	}

	// Wait for all the workers in the WaitGroup to finish.
	println("Waiting for workers to finish")
	wg.Wait()
	println("Workers finished")
	close(results)

	// Optionally, collect all the results.
	var maps []categoryMap
	for r := range results {
		maps = append(maps, r)
	}
	// fmt.Println("LOOK HERE:", maps[0].sourceToDestination[50])

	// Sorting the maps by the lowest ID to the highest ID
	var sortedMaps []categoryMap
	for i := 1; i <= 7; i++ {
		for _, m := range maps {
			if m.ID == i {
				sortedMaps = append(sortedMaps, m)
			}
		}
	}
	// println(sortedMaps[0].mapName)

	// Finding the locaiton for each seed, by tracking through the seed-to-soil map, soil-to-fertilizer map, fertilizer-to-water map,
	// water-to-light map, light-to-temperature map, temperature-to-humidity map, and humidity-to-location map.
	// But if the map does not contain a value for the key it is given, then value = key
	// First we make a list of each seed's location
	var seedLocations []int = make([]int, len(seeds))
	// fmt.Println("Seeds:", seeds)

	println("Finding the location for each seed")
	for seedID, seed := range seeds {
		var currentValue int = seed
		for i, m := range sortedMaps {
			var newValue, ok = m.sourceToDestination[currentValue]
			// Checking that we're checking the right map, otherwise print an error
			if m.ID != i+1 {
				fmt.Println("Error - Wrong map ID:", m.ID, "Expected ID:", i+1)
			}

			if !ok {
				// If the map does not contain the key, then the value is the key
				// fmt.Println("Map", m.mapName, "does not contain the key", currentValue)
				continue
			} else {
				// fmt.Println("Map", m.mapName, "contains the key", currentValue)
				currentValue = newValue
			}
		}
		seedLocations[seedID] = currentValue
	}
	// fmt.Println("Seed:", seeds[0], "Last location", seedLocations[0])

	var lowestLocation int = math.MaxInt32
	for _, seedLocation := range seedLocations {
		if seedLocation < lowestLocation {
			lowestLocation = seedLocation
		}
	}
	fmt.Println("Lowest location:", lowestLocation)

}

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

/* Sorting
ID 1 = seed-to-soil map:
ID 2 = soil-to-fertilizer map:
ID 3 = fertilizer-to-water map:
ID 4 = water-to-light map:
ID 5 = light-to-temperature map:
ID 6 = temperature-to-humidity map:
ID 7 = humidity-to-location map:
*/
