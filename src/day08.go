package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	currentID string
	leftID    string
	rightID   string
}

func parseNodes(str []string) []node {
	var nodes []node
	for _, line := range str {
		var newNode node = node{currentID: line[0:3], leftID: line[7:10], rightID: line[12:15]}
		nodes = append(nodes, newNode)
	}
	return nodes
}

func main() {
	file, err := os.Open("files/day08")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		var line string = scanner.Text()
		lines = append(lines, line)
	}

	var instructions string = lines[0]
	var nodes []node = parseNodes(lines[2:])

	// Creating a hashmap to store the nodes, where the key is the ID
	var nodesMap = make(map[string]node)
	for _, node := range nodes {
		nodesMap[node.currentID] = node
	}
	var currentNode node = nodesMap["AAA"]

	var steps int = 0
	for {
		if currentNode.currentID == "ZZZ" {
			fmt.Println("Node ZZZ reached in", steps, "steps")
			break
		} else {
			switch instructions[steps%len(instructions)] {
			case 'L':
				currentNode = nodesMap[currentNode.leftID]
			case 'R':
				currentNode = nodesMap[currentNode.rightID]
			default:
				fmt.Println("Invalid instruction", instructions[steps])
			}

			steps++
		}
	}
}
