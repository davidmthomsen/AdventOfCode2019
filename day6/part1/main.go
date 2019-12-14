package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput(path string) map[string][]string {
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	orbitMap := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		splittedLine := strings.Split(line, ")")
		object := splittedLine[0]
		orbitingObject := splittedLine[1]

		if orbitingObjects, ok := orbitMap[object]; ok {
			orbitingObjects = append(orbitingObjects, orbitingObject)
			orbitMap[object] = orbitingObjects
		} else {
			var orbitingObjects []string
			orbitingObjects = append(orbitingObjects, orbitingObject)
			orbitMap[object] = orbitingObjects
		}
	}

	err = file.Close()

	if err != nil {
		log.Fatalf("failed closing file: %s", err)
	}

	return orbitMap
}

func exploreMap(orbitMap map[string][]string, currentObject string, depth int) int {
	orbitingObjects := orbitMap[currentObject]

	exploredDepth := 0

	for _, orbitingObject := range orbitingObjects {
		exploredDepth += exploreMap(orbitMap, orbitingObject, depth+1)
	}

	return exploredDepth + depth
}

func main() {
	orbitMap := readInput("input.txt")
	totalOrbits := exploreMap(orbitMap, "COM", 0)
	fmt.Println(totalOrbits)
}
