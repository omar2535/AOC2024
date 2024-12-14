package day13

import (
	"aoc2024/internal"
	"fmt"
	"strings"
)

type Machine struct {
	buttonA Button
	buttonB Button
	prize   Position
}

type Position struct {
	x int
	y int
}

type Button struct {
	x int
	y int
}

func PartOne(isTest bool) {
	// Some initial variables
	var fileContents []string

	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day13/day13_example.txt")
	} else {
		fileContents = internal.ReadFileIntoArray("res/day13/day13.txt")
	}

	// Split into array of Machines
	machines := make([]Machine, 0)
	for i := 0; i < len(fileContents); i += 4 {
		buttonA := getButtonFromString(fileContents[i])
		buttonB := getButtonFromString(fileContents[i+1])
		prize := getPrizeFromString(fileContents[i+2])
		machines = append(machines, Machine{buttonA, buttonB, prize})
	}

	// Go through each machine and get the minimum way to press the buttons
	tokens := 0
	for _, machine := range machines {
		fmt.Println(getMinUsingSystemOfEquation(machine))
		tokens += getMinUsingSystemOfEquation(machine)
	}

	fmt.Println("Num tokens used:", tokens)
}

func PartTwo(isTest bool) {
	// Some initial variables
	var fileContents []string

	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day13/day13_example.txt")
	} else {
		fileContents = internal.ReadFileIntoArray("res/day13/day13.txt")
	}

	// Split into array of Machines
	machines := make([]Machine, 0)
	for i := 0; i < len(fileContents); i += 4 {
		buttonA := getButtonFromString(fileContents[i])
		buttonB := getButtonFromString(fileContents[i+1])
		prize := getPrizeFromString(fileContents[i+2])
		prize.x += 10000000000000
		prize.y += 10000000000000
		machines = append(machines, Machine{buttonA, buttonB, prize})
	}

	// Go through each machine and get the minimum way to press the buttons
	tokens := 0
	for _, machine := range machines {
		fmt.Println(getMinUsingSystemOfEquation(machine))
		tokens += getMinUsingSystemOfEquation(machine)
	}

	fmt.Println("Num tokens used:", tokens)
}
func getMinUsingSystemOfEquation(machine Machine) int {
	// First, get a single solution for the equations
	a1 := machine.buttonA.x
	b1 := machine.buttonB.x
	c1 := machine.prize.x
	a2 := machine.buttonA.y
	b2 := machine.buttonB.y
	c2 := machine.prize.y

	// Check if the lines are parallel, if they are, then there is no solution
	if a1*b2 == a2*b1 {
		return 0
	}

	// x is the number of buttonA presses
	// y is the number of buttonB presses
	x := float64(float64(b2*c1-b1*c2) / float64(a1*b2-a2*b1))
	y := float64(float64(a1*c2-a2*c1) / float64(a1*b2-a2*b1))

	// Check if x and y are integers (we need button presses that aren't fractions)
	if x != float64(int(x)) || y != float64(int(y)) {
		return 0
	}

	return int(x)*3 + int(y)
}

// Keep track of 2 main things:
// For any score (x, y),
// - all the buttons that can be pressed to get to that score
func getMinPressesForMachine(machine Machine) int {
	var pressesAtPosition map[Position][]map[Button]int = make(map[Position][]map[Button]int)
	var targetX = machine.prize.x
	var targetY = machine.prize.y
	buttonA := machine.buttonA
	buttonB := machine.buttonB

	// Do the buttonA presses first
	for _, currentButton := range []Button{buttonA, buttonB} {
		for x := 0; x <= targetX; x++ {
			for y := 0; y <= targetY; y++ {
				fmt.Println(x, y)
				if x-currentButton.x == 0 && y-currentButton.y == 0 {
					pressesAtPosition[Position{x, y}] = make([]map[Button]int, 0)
					pressesAtPositionMap := make(map[Button]int)
					pressesAtPositionMap[currentButton] = 1
					pressesAtPosition[Position{x, y}] = append(pressesAtPosition[Position{x, y}], pressesAtPositionMap)
				} else if x-currentButton.x >= 0 && y-currentButton.y >= 0 {
					if pressesAtPosition[Position{x - currentButton.x, y - currentButton.y}] != nil {
						// Create the array at the location if it doesn't exist
						if pressesAtPosition[Position{x, y}] == nil {
							pressesAtPosition[Position{x, y}] = make([]map[Button]int, 0)
						}
						// Go through each previous press combination and add the current button press
						previousPresses := pressesAtPosition[Position{x - currentButton.x, y - currentButton.y}]
						for _, previousPress := range previousPresses {
							pressesAtPositionMap := make(map[Button]int)
							for button, presses := range previousPress {
								pressesAtPositionMap[button] = presses
							}
							pressesAtPositionMap[currentButton] = pressesAtPositionMap[currentButton] + 1
							pressesAtPosition[Position{x, y}] = append(pressesAtPosition[Position{x, y}], pressesAtPositionMap)
						}
					}
				}
			}
		}
	}

	// Now do the buttonB presses
	return 0
}

func getMinPressesForMachineRecursive(machine Machine, remainingX int, remainingY int, numButtonAPresses int, numButtonBPrseses int, depth int) int {
	if remainingX == 0 && remainingY == 0 {
		return depth
	} else if remainingX < 0 || remainingY < 0 {
		return 1000000000000
	} else if numButtonAPresses > 100 || numButtonBPrseses > 100 {
		return 1000000000000
	} else {
		return min(
			getMinPressesForMachineRecursive(machine, remainingX-machine.buttonA.x, remainingY-machine.buttonA.y, numButtonAPresses+1, numButtonBPrseses, depth+1),
			getMinPressesForMachineRecursive(machine, remainingX-machine.buttonB.x, remainingY-machine.buttonB.y, numButtonAPresses, numButtonBPrseses+1, depth+1),
		)
	}
}

func getButtonFromString(buttonStr string) Button {
	xyString := strings.Split(buttonStr, ":")[1]
	xString := strings.Split(xyString, ", ")[0]
	yString := strings.Split(xyString, ", ")[1]

	x := internal.GetNumFromString(strings.Split(xString, "+")[1])
	y := internal.GetNumFromString(strings.Split(yString, "+")[1])
	return Button{x, y}
}

func getPrizeFromString(prizeStr string) Position {
	xyString := strings.Split(prizeStr, ":")[1]
	xString := strings.Split(xyString, ", ")[0]
	yString := strings.Split(xyString, ", ")[1]

	x := internal.GetNumFromString(strings.Split(xString, "=")[1])
	y := internal.GetNumFromString(strings.Split(yString, "=")[1])
	return Position{x, y}
}
