package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const expOutput = 19690720

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

func tryNounVerb(input []int, noun, verb int) int {

	inputCopy := make([]int, len(input))
	copy(inputCopy, input)

	inputCopy[1] = noun
	inputCopy[2] = verb

	instructionIndex := 0
	for inputCopy[instructionIndex] != 99 {
		indexv1 := inputCopy[instructionIndex+1]
		indexv2 := inputCopy[instructionIndex+2]
		indexv3 := inputCopy[instructionIndex+3]

		switch inputCopy[instructionIndex] {
		case 1:
			inputCopy[indexv3] = inputCopy[indexv1] + inputCopy[indexv2]
		case 2:
			inputCopy[indexv3] = inputCopy[indexv1] * inputCopy[indexv2]
		default:
			log.Fatalf("[*] Unknown opcode %d at %d\n", inputCopy[instructionIndex], instructionIndex)
		}
		instructionIndex += 4
	}
	return inputCopy[0]
}

func main() {
	input := readInput("input.txt")

loop:
	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			output := tryNounVerb(input, noun, verb)
			if output == expOutput {
				fmt.Println(noun*100 + verb)
				break loop
			}
		}
	}
}
