package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type parameterMode int
type opcodeType int

const (
	positionMode  parameterMode = 0
	immediateMode parameterMode = 1

	addCode    opcodeType = 1
	multiCode  opcodeType = 2
	inputCode  opcodeType = 3
	outputCode opcodeType = 4
	stopCode   opcodeType = 99
)

func readInput(path string) []int {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("[*] Could not read file: %s\n", err)
	}

	var input []int
	splitData := strings.Split(string(data), ",")

	for _, valueStr := range splitData {
		valueInt, err := strconv.Atoi(valueStr)
		if err != nil {
			log.Fatalf("[*] Could not convert to int: %s\n", err)
		}
		input = append(input, valueInt)
	}
	return input
}

func parseInstruction(instruction int) (opcodeType, []parameterMode) {
	var paramsMode []parameterMode

	opcode := opcodeType(instruction % 100)
	instruction /= 100

	for instruction > 0 {
		paramMode := parameterMode(instruction % 10)
		paramsMode = append(paramsMode, paramMode)

		instruction /= 10
	}

	return opcode, paramsMode
}

func readValue(input []int, instructionIndex, parameterIndex int, parametersMode []parameterMode) int {
	paramMode := positionMode
	if len(parametersMode) > parameterIndex-1 {
		paramMode = parametersMode[parameterIndex-1]
	}

	var output int

	switch paramMode {
	case positionMode:
		readIndex := input[instructionIndex+parameterIndex]
		output = input[readIndex]
	case immediateMode:
		output = input[instructionIndex+parameterIndex]
	default:
		log.Fatalf("[*] Unknown parameter mode: %d from instruction at index: %d\n", paramMode, instructionIndex)
	}
	return output
}

func writeValue(input []int, index, value int) {
	readIndex := input[index]
	input[readIndex] = value
}

func main() {
	input := readInput("input.txt")

	userInput := 1

	instructionIndex := 0
instructionLoop:
	for {
		opcode, paramsMode := parseInstruction(input[instructionIndex])

		switch opcode {
		case addCode:
			v1 := readValue(input, instructionIndex, 1, paramsMode)
			v2 := readValue(input, instructionIndex, 2, paramsMode)
			writeValue(input, instructionIndex+3, v1+v2)
			instructionIndex += 4
		case multiCode:
			v1 := readValue(input, instructionIndex, 1, paramsMode)
			v2 := readValue(input, instructionIndex, 2, paramsMode)
			writeValue(input, instructionIndex+3, v1*v2)
			instructionIndex += 4
		case inputCode:
			writeValue(input, instructionIndex+1, userInput)
			instructionIndex += 2
		case outputCode:
			v1 := readValue(input, instructionIndex, 1, paramsMode)
			fmt.Println(v1)
			instructionIndex += 2
		case stopCode:
			break instructionLoop
		default:
			log.Fatalf("[*] Unknow opcode: %d from instruction at index: %d\n", opcode, instructionIndex)
		}
	}
}
