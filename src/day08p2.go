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

func endsWith(suffix rune, ID string) bool {
	return rune(ID[len(ID)-1]) == suffix
}

func doAllEndWithZ(nodes []node) bool {
	for _, node := range nodes {
		if !endsWith('Z', node.currentID) {
			return false
		}
	}
	return true
}

func moveAllTo(side rune, nodes []node, nodeHash map[string]node) []node {
	for i, _ := range nodes {
		switch side {
		case 'L':
			nodes[i] = nodeHash[nodes[i].leftID]
		case 'R':
			nodes[i] = nodeHash[nodes[i].rightID]
		default:
			fmt.Println("Invalid side", side)
		}
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

	// currentNodes is initialized with all the nodes that end with A
	var currentNodes []node
	for _, node := range nodes {
		if endsWith('A', node.currentID) {
			currentNodes = append(currentNodes, node)
		}
	}

	var steps int = 0
	for {
		if doAllEndWithZ(currentNodes) {
			fmt.Println("All current node IDs end with a Z, stopped at step", steps)
			break
		} else {
			var nextSide = instructions[steps%len(instructions)]
			currentNodes = moveAllTo(rune(nextSide), currentNodes, nodesMap)
			steps++
		}
	}
}
