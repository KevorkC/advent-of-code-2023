package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type raceInfo struct {
	time           int
	recordDistance int
}

func fromListsToRaceInfo(timeList []int, recordList []int) []raceInfo {
	var pairList []raceInfo
	for i := 0; i < len(timeList); i++ {
		newPair := raceInfo{time: timeList[i], recordDistance: recordList[i]}
		pairList = append(pairList, newPair)
	}
	return pairList
}

func parseStringToIntList(input string) []int {
	var intList []int
	var fields []string = strings.Fields(input)
	for _, subStrNumber := range fields {
		var number, e = strconv.Atoi(subStrNumber)
		if e != nil {
			fmt.Print("Error converting substring to number:", e)
		} else {
			intList = append(intList, number)
		}
	}
	return intList
}

func calculatePossibleSolutions(race raceInfo) int {
	// The solutions are all the options that are faster than recordDistance
	var possibleSolutions int = 0
	for windUp := 0; windUp < race.time; windUp++ {
		var restOfRace int = race.time - windUp
		var distance int = windUp * restOfRace
		if distance > race.recordDistance {
			possibleSolutions++
		}
	}
	return possibleSolutions
}

func concatenateInts(ints []int) int {
	var resultStr string
	for _, num := range ints {
		resultStr += strconv.Itoa(num)
	}

	var resultInt, err = strconv.Atoi(resultStr)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}
	return resultInt
}

func main() {
	file, err := os.Open("files/day06")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileLines []string

	for scanner.Scan() {
		var line string = scanner.Text()
		fileLines = append(fileLines, line)
	}

	var raceTimeList []int = []int{concatenateInts(parseStringToIntList(fileLines[0][9:]))}
	var raceRecordList []int = []int{concatenateInts(parseStringToIntList(fileLines[1][9:]))}

	var raceList []raceInfo = fromListsToRaceInfo(raceTimeList, raceRecordList)

	fmt.Println("raceList", raceList)

	var possibleSolutionsMultiplied int = 1
	for _, race := range raceList {
		possibleSolutionsMultiplied *= calculatePossibleSolutions(race)
	}

	fmt.Println(possibleSolutionsMultiplied)
}
