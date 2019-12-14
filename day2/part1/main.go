package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readInput(path string) []int {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("[*] Could not read file: %s\n", err)
	}

	var input []int

	splitData := strings.Split(string(data), ",")

	for _, valStr := range splitData {
		valInt, err := strconv.Atoi(valStr)
		if err != nil {
			log.Fatalf("[*] Could not convert to int: %s\n", err)
		}
		input = append(input, valInt)
	}
	return input
}

func main() {

	input := readInput("input.txt")

	input[1] = 12
	input[2] = 2

	instructionIndex := 0
	for input[instructionIndex] != 99 {
		indexv1 := input[instructionIndex+1]
		indexv2 := input[instructionIndex+2]
		indexv3 := input[instructionIndex+3]

		switch input[instructionIndex] {
		case 1:
			input[indexv3] = input[indexv1] + input[indexv2]
		case 2:
			input[indexv3] = input[indexv1] * input[indexv2]
		default:
			log.Fatalf("[*] Unknown opcode %d at %d\n", input[instructionIndex], instructionIndex)
		}
		instructionIndex += 4
	}
	fmt.Println(input[0])

}
