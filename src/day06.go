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

	var raceTimeList []int = parseStringToIntList(fileLines[0][9:])
	var raceRecordList []int = parseStringToIntList(fileLines[1][9:])

	var raceList []raceInfo = fromListsToRaceInfo(raceTimeList, raceRecordList)

	fmt.Println("raceList", raceList)
}
