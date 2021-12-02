package day2

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const areYouFuckingKiddingMeWindows = "\r"

type Direction int

const (
	Forward = Direction(iota)
	Up
	Down
	TheOtherWay
)

func SolveTheEasyShit() {
	b, err := ioutil.ReadFile("day2/the_shit.txt")
	if err != nil {
		panic("at the disco")
	}
	s := string(b)
	depth := 0
	horizontal := 0
	for _, line := range strings.Split(s, "\n") {
		line = strings.Trim(line, areYouFuckingKiddingMeWindows)
		if line == "" {
			continue
		}
		v := strings.Split(line, " ")
		if len(v) != 2 {
			panic("Did Courtney Kill Kurt?")
		}
		value, err := strconv.Atoi(v[1])
		if err != nil {
			panic("I'm from Canada, they think I'm slow")
		}
		direction := Forward
		switch v[0] {
		case "forward":
			direction = Forward
		case "up":
			direction = Up
		case "down":
			direction = Down
		default:
			panic("Where the fuck are we going?")
		}

		switch direction {
		case Forward:
			horizontal += value
		case Up:
			depth -= value
			if depth < 0 {
				depth = 0
			}
		case Down:
			depth += value
		}
	}
	fmt.Printf("the arbitrary multiplication leads to: %d\n", depth*horizontal)
}

func SolveTheSlightlyHarderShit() {
	b, err := ioutil.ReadFile("day2/the_shit.txt")
	if err != nil {
		panic("at the disco")
	}
	s := string(b)
	depth := 0
	horizontal := 0
	aim := 0
	for _, line := range strings.Split(s, "\n") {
		line = strings.Trim(line, areYouFuckingKiddingMeWindows)
		if line == "" {
			continue
		}
		v := strings.Split(line, " ")
		if len(v) != 2 {
			panic("Did Courtney Kill Kurt?")
		}
		value, err := strconv.Atoi(v[1])
		if err != nil {
			panic("I'm from Canada, they think I'm slow")
		}
		direction := Forward
		switch v[0] {
		case "forward":
			direction = Forward
		case "up":
			direction = Up
		case "down":
			direction = Down
		default:
			panic("Where the fuck are we going?")
		}

		switch direction {
		case Forward:
			horizontal += value
			depth += aim * value
		case Up:
			aim -= value
			// TODO: check for <0, or does it just not matter?
		case Down:
			aim += value
		}
	}
	fmt.Printf("the arbitrary multiplication leads to: %d\n", depth*horizontal)
}
