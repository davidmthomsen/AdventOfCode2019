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

	addCode         opcodeType = 1
	mulCode         opcodeType = 2
	inputCode       opcodeType = 3
	outputCode      opcodeType = 4
	jumpIfTrueCode  opcodeType = 5
	jumpIfFalseCode opcodeType = 6
	lessThanCode    opcodeType = 7
	equalsCode      opcodeType = 8

	stopCode opcodeType = 99
)

func readInput(path string) []int {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("[*] Could not read file :%s\n", err)
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

func readValue(input []int, instructionPointer, parameterIndex int, parametersMode []parameterMode) int {

	paramMode := positionMode
	if len(parametersMode) > parameterIndex-1 {
		paramMode = parametersMode[parameterIndex-1]
	}

	var output int
	switch paramMode {
	case positionMode:
		readIndex := input[instructionPointer+parameterIndex]
		output = input[readIndex]
	case immediateMode:
		output = input[instructionPointer+parameterIndex]
	default:
		log.Fatalf("[*] Unknown parameter mode; %d from instruction at index: %d\n", paramMode, instructionPointer)
	}

	return output
}

func writeValue(input []int, index, value int) {
	readIndex := input[index]
	input[readIndex] = value
}

func main() {
	input := readInput("input.txt")

	userInput := 5

	instructionPointer := 0

instructionLoop:
	for {
		opcode, paramsMode := parseInstruction(input[instructionPointer])

		switch opcode {
		case addCode:
			v1 := readValue(input, instructionPointer, 1, paramsMode)
			v2 := readValue(input, instructionPointer, 2, paramsMode)
			writeValue(input, instructionPointer+3, v1+v2)
			instructionPointer += 4
		case mulCode:
			v1 := readValue(input, instructionPointer, 1, paramsMode)
			v2 := readValue(input, instructionPointer, 2, paramsMode)
			writeValue(input, instructionPointer+3, v1*v2)
			instructionPointer += 4
		case inputCode:
			writeValue(input, instructionPointer+1, userInput)
			instructionPointer += 2
		case outputCode:
			v1 := readValue(input, instructionPointer, 1, paramsMode)
			fmt.Println(v1)
			instructionPointer += 2
		case jumpIfTrueCode:
			v1 := readValue(input, instructionPointer, 1, paramsMode)
			v2 := readValue(input, instructionPointer, 2, paramsMode)
			if v1 != 0 {
				instructionPointer = v2
			} else {
				instructionPointer += 3
			}
		case jumpIfFalseCode:
			v1 := readValue(input, instructionPointer, 1, paramsMode)
			v2 := readValue(input, instructionPointer, 2, paramsMode)
			if v1 == 0 {
				instructionPointer = v2
			} else {
				instructionPointer += 3
			}
		case lessThanCode:
			v1 := readValue(input, instructionPointer, 1, paramsMode)
			v2 := readValue(input, instructionPointer, 2, paramsMode)
			valueToWrite := 0
			if v1 < v2 {
				valueToWrite = 1
			}

			writeValue(input, instructionPointer+3, valueToWrite)
			instructionPointer += 4
		case equalsCode:
			v1 := readValue(input, instructionPointer, 1, paramsMode)
			v2 := readValue(input, instructionPointer, 2, paramsMode)
			valueToWrite := 0
			if v1 == v2 {
				valueToWrite = 1
			}
			writeValue(input, instructionPointer+3, valueToWrite)
			instructionPointer += 4
		case stopCode:
			break instructionLoop
		default:
			log.Fatalf("[*] Unknown opcode: %d from instruction at index: %d\n", opcode, instructionPointer)
		}
	}
}
