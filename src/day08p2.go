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

type nodeWithZ struct {
	foundIn int
	nodeID  string
}

// Function to find the Greatest Common Divisor (GCD)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to find the Least Common Multiple (LCM) of two numbers
func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a / gcd(a, b) * b // Rearranged to avoid potential overflow
}

// Function to find the LCM of a list of integers
func lcmOfList(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	result := numbers[0]
	for _, number := range numbers[1:] {
		result = lcm(result, number)
		if result == 0 {
			break
		}
	}
	return result
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

func areAllZNodesFound(nodeWithZ []nodeWithZ) bool {
	for _, node := range nodeWithZ {
		if node.foundIn == -1 {
			return false
		}
	}
	return true
}

func calculateTotalSteps(targetNodes []nodeWithZ) int {
	if !areAllZNodesFound(targetNodes) {
		fmt.Println("Error: calculateTotalSteps called before all Z nodes are found")
	}

	// Making the list of integers that will be used to find the LCM
	var integersToLCM []int
	for _, node := range targetNodes {
		integersToLCM = append(integersToLCM, node.foundIn)
	}

	fmt.Println("Integers to LCM:", integersToLCM)
	// Using Least Common Multiple
	return lcmOfList(integersToLCM)
}

func findNewZNodes(currentNodes []node, targetNodes []nodeWithZ, currentStep int) []nodeWithZ {
	for _, currentNode := range currentNodes {
		for i, _ := range targetNodes {
			if currentNode.currentID == targetNodes[i].nodeID {
				if targetNodes[i].foundIn == -1 {
					fmt.Println("Found new Z node:", currentNode.currentID, "at step:", currentStep)
					targetNodes[i].foundIn = currentStep
				}
			}
		}
	}
	return targetNodes
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

	// zNodes are all the nodes that end with Z
	var zNodes []node
	for _, node := range nodes {
		if endsWith('Z', node.currentID) {
			zNodes = append(zNodes, node)
		}
	}

	// targetNodes are all the the zNodes that are found in the instructions
	var targetNodes []nodeWithZ
	for _, node := range zNodes {
		targetNodes = append(targetNodes, nodeWithZ{foundIn: -1, nodeID: node.currentID})
	}

	var steps int = 0
	for {
		if areAllZNodesFound(targetNodes) {
			fmt.Println("Target Nodes:", targetNodes)
			fmt.Println("Found all loops in", calculateTotalSteps(targetNodes), "steps")
			break
		}

		var nextSide = instructions[steps%len(instructions)]
		currentNodes = moveAllTo(rune(nextSide), currentNodes, nodesMap)
		steps++
		targetNodes = findNewZNodes(currentNodes, targetNodes, steps)

		if steps%10000 == 0 {
			fmt.Println("Currently at step", steps)
		}
	}
}
