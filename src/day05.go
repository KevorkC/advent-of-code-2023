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
	ID              int
	mapName         string
	destinationList []int
	sourceList      []int
	rangeLengthList []int
	// sourceToDestination map[int]int
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

		// sdm.sourceToDestination = make(map[int]int)
		var sourceList []int
		var destinationList []int
		var rangeLengthList []int
		for _, line := range job[1:] {
			for i, subStrNumber := range strings.Fields(line) {
				if i == 0 {
					var destination, e = strconv.Atoi(subStrNumber)
					if e != nil {
						fmt.Print("Error converting destination to number:", e)
					} else {
						destinationList = append(destinationList, destination)
					}
				} else if i == 1 {
					var source, e = strconv.Atoi(subStrNumber)
					if e != nil {
						fmt.Print("Error converting source to number:", e)
					} else {
						sourceList = append(sourceList, source)
					}
				} else if i == 2 {
					var rangeLength, e = strconv.Atoi(subStrNumber)
					if e != nil {
						fmt.Print("Error converting rangeLength to number:", e)
					} else {
						rangeLengthList = append(rangeLengthList, rangeLength)
					}
				}
			}

		}

		sdm.destinationList = destinationList
		sdm.sourceList = sourceList
		sdm.rangeLengthList = rangeLengthList
		results <- sdm
	}
}

// The following function calculates the destination from the source, given the map
func fromSourceToDestination(sourceTarget int, category categoryMap) int {
	var sourceList []int = category.sourceList
	var destinationList []int = category.destinationList
	var rangeLengthList []int = category.rangeLengthList

	for i := 0; i < len(sourceList); i++ {
		sourceStart := sourceList[i]
		destinationStart := destinationList[i]
		rangeLength := rangeLengthList[i]

		// Check if the sourceTarget is within the current range
		if sourceTarget >= sourceStart && sourceTarget < sourceStart+rangeLength {
			// Calculate the offset from the start of the source range
			offset := sourceTarget - sourceStart
			// Return the corresponding destination number
			return destinationStart + offset
		}
	}
	// If the sourceTarget is not in any range, return the sourceTarget itself
	return sourceTarget
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

	// Sorting the maps by the lowest ID to the highest ID
	var sortedMaps []categoryMap
	for i := 1; i <= 7; i++ {
		for _, m := range maps {
			if m.ID == i {
				sortedMaps = append(sortedMaps, m)
			}
		}
	}

	// First we make a list of each seed's location
	var seedLocations []int = make([]int, len(seeds))
	// fmt.Println("Seeds:", seeds)

	println("Finding the location for each seed")
	for seedID, seed := range seeds {
		var currentValue int = seed
		for i, m := range sortedMaps {
			var destination = fromSourceToDestination(currentValue, m)
			currentValue = destination
			// Checking that we're checking the right map, otherwise print an error
			if m.ID != i+1 {
				fmt.Println("Error - Wrong map ID:", m.ID, "Expected ID:", i+1)
			}
		}
		seedLocations[seedID] = currentValue
	}

	var lowestLocation int = math.MaxInt32
	for _, seedLocation := range seedLocations {
		if seedLocation < lowestLocation {
			lowestLocation = seedLocation
		}
	}
	fmt.Println("Lowest location:", lowestLocation)

}
