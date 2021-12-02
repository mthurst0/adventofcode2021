package day1

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func errShouldHaveAStringFunc(err error) string {
	return err.Error()
}

const areYouFuckingKiddingMeWindows = "\r"

func SolveTheEasyShit() {
	data, err := ioutil.ReadFile("day1/the_shit.txt")
	if err != nil {
		panic(err)
	}
	s := string(data)
	isThisTheFirstFuckingOne := true
	last := 0
	countToWinTheChallenge := 0
	for _, line := range strings.Split(s, "\n") {
		line = strings.Trim(line, areYouFuckingKiddingMeWindows)
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			panic("in the streets of london: " + errShouldHaveAStringFunc(err))
		}
		if num > last {
			if !isThisTheFirstFuckingOne {
				countToWinTheChallenge++
			}
		}
		isThisTheFirstFuckingOne = false
		last = num
	}
	fmt.Printf("the answer is 42 and/or: %d\n", countToWinTheChallenge)
}

func turnt(bunchOfData string) []int {
	var result []int
	for _, line := range strings.Split(bunchOfData, "\n") {
		line = strings.Trim(line, areYouFuckingKiddingMeWindows)
		if line == "" {
			continue
		}
		theJuice, err := strconv.Atoi(line)
		if err != nil {
			panic("Bruges is a shit hole")
		}
		result = append(result, theJuice)
	}
	return result
}

func SolveTheSlightlyHarderShit() {
	data, err := ioutil.ReadFile("day1/the_shit.txt")
	if err != nil {
		panic(err)
	}
	allTheShit := turnt(string(data))
	if len(allTheShit) <= 3 {
		panic("why you ask for the impossible?")
	}
	cur0 := allTheShit[0]
	cur1 := allTheShit[1]
	cur2 := allTheShit[2]
	ourMakerNeedsTheirAnswer := 0
	for i := 3; i < len(allTheShit); i++ {
		v := cur1 + cur2 + allTheShit[i]
		if v > (cur0 + cur1 + cur2) {
			ourMakerNeedsTheirAnswer++
		}
		cur0 = cur1
		cur1 = cur2
		cur2 = allTheShit[i]
	}
	fmt.Printf("yes, we have it (maybe?): %d\n", ourMakerNeedsTheirAnswer)
}
