package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func stringToInt(inputString string) int {
	stringToInt, err := strconv.Atoi(inputString)
	if err == nil {
		return stringToInt
	}
	return 0
}

func findIntInLine(line []byte) int {
	var firstVal, secondVal string
	lineByteToString := string(line)
	regexMatchMultiInt := regexp.MustCompile("[0-9].*[0-9]").FindString(lineByteToString)
	regexMatchSingleInt := regexp.MustCompile("[0-9]").FindString(lineByteToString)
	if len(regexMatchMultiInt) > 0 {
		firstVal = string(regexMatchMultiInt[0])
		secondVal = string(regexMatchMultiInt[len(regexMatchMultiInt)-1])
	} else if len(regexMatchSingleInt) > 0 {
		firstVal, secondVal = string(regexMatchSingleInt[0]), string(regexMatchSingleInt[0])
	}
	return stringToInt(fmt.Sprintf("%s%s", firstVal, secondVal))
}

func main() {

	var intFinds []int
	var finalSum int
	readFile, err := os.Open("./input/day1/input.txt")
	if err == nil {
		fileScanner := bufio.NewScanner(readFile)
		for fileScanner.Scan() {
			intFinds = append(intFinds, findIntInLine(fileScanner.Bytes()))
		}
		for _, integer := range intFinds {
			finalSum += integer
		}
		fmt.Println(finalSum)
		readFile.Close()
	}
}
