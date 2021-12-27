package day17

import (
	"fmt"
	"testing"
)

type Puzzle struct {
	XTargetMin int
	XTargetMax int
	YTargetMin int
	YTargetMax int
}

func trickShot(p Puzzle, startXVel, startYVel int) (bool, int) {
	var xPos, yPos int
	xVel := startXVel
	yVel := startYVel
	maxYReached := 0
	for {
		xPos += xVel
		yPos += yVel
		if yPos > maxYReached {
			maxYReached = yPos
		}
		if xVel > 0 {
			xVel--
		}
		yVel--
		if xPos > p.XTargetMax || yPos < p.YTargetMin {
			return false, 0
		}
		if xPos >= p.XTargetMin && xPos <= p.XTargetMax && yPos >= p.YTargetMin && yPos <= p.YTargetMax {
			return true, maxYReached
		}
	}
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func TestTrickShot(t *testing.T) {
	// sample puzzle
	//samplePuzzle := Puzzle{
	// XTargetMin: 20,
	// XTargetMax: 30,
	// YTargetMin: -10,
	// YTargetMax: -5,
	//}

	// easy puzzle
	easyPuzzle := Puzzle{
		XTargetMin: 155,
		XTargetMax: 215,
		YTargetMin: -132,
		YTargetMax: -72,
	}

	soln := 0
	solnCount := 0
	//p := samplePuzzle
	p := easyPuzzle
	for x := 1; x <= p.XTargetMax; x++ {
		for y := p.YTargetMin; y < absInt(p.YTargetMin)*2; y++ {
			startXVel := x
			startYVel := y
			solutionFound, maxYReached := trickShot(p, startXVel, startYVel)
			if solutionFound {
				fmt.Printf("found solution: %d,%d reached: %d\n", startXVel, startYVel, maxYReached)
				solnCount++
				if maxYReached > soln {
					soln = maxYReached
				}
			} else {
				//fmt.Printf("no solution found for %d,%d", startXVel, startYVel)
			}

		}
	}
	fmt.Printf("soln=%d\n", soln)
	fmt.Printf("solnCount=%d\n", solnCount)
}
