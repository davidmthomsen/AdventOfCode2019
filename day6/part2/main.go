package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type objectOrbit struct {
	orbiting  string
	orbitedBy []string
}

func readInput(path string) map[string]objectOrbit {
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("[*] Failed opening file: %s", err)
	}

	orbitMap := make(map[string]objectOrbit)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		scanner.Text()

		line := scanner.Text()
		splitLine := strings.Split(line, ")")
		object := splitLine[0]
		orbitingObj := splitLine[1]

		objectOrbit := orbitMap[object]
		objectOrbit.orbitedBy = append(objectOrbit.orbitedBy, orbitingObj)
		orbitMap[object] = objectOrbit

		orbitingObjectOrbit := orbitMap[orbitingObj]
		orbitingObjectOrbit.orbiting = object
		orbitMap[orbitingObj] = orbitingObjectOrbit

	}

	err = file.Close()

	if err != nil {
		log.Fatalf("[*] Failed closing file: %s", err)
	}

	return orbitMap
}

func exploreMap(orbitMap map[string]objectOrbit, fromObject, currentObject, stopObject string) int {
	var exploredDepth int

	if currentObject == stopObject {
		return 1
	}

	currentObjectOrbits := orbitMap[currentObject]

	if currentObjectOrbits.orbiting != fromObject {
		exploredDepth += exploreMap(orbitMap, currentObject, currentObjectOrbits.orbiting, stopObject)
		if exploredDepth > 0 {
			return exploredDepth + 1
		}
	}

	for _, orbitingObject := range currentObjectOrbits.orbitedBy {
		if orbitingObject == fromObject {
			continue
		}

		exploredDepth += exploreMap(orbitMap, currentObject, orbitingObject, stopObject)

		if exploredDepth > 0 {
			return exploredDepth + 1
		}
	}
	return 0
}

func main() {
	orbitMap := readInput("input.txt")

	startObject := orbitMap["YOU"].orbiting
	stopObject := orbitMap["SAN"].orbiting

	minPath := exploreMap(orbitMap, "YOU", startObject, stopObject) - 1

	fmt.Println(minPath)

}
