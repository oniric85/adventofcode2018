package main

import (
	"fmt"
	"log"
	"sort"
)

func FindReadySteps(instructions map[string][]string) (ready []string) {
	for step, before := range instructions {
		// a step is ready if the length of the associated array of steps is zero
		// this means that no step is needed before it can be carried on
		if len(before) == 0 {
			ready = append(ready, step)
		}
	}

	// return the list sorted to cope with puzzle constraints
	sort.Strings(ready)

	return ready
}

func FindOrder(instructions map[string][]string) (ordered string) {
	ready := FindReadySteps(instructions)
	var step string

	for len(ready) > 0 {
		step, ready = Shift(ready)

		ordered += step

		// now we need to process the map and update the list of ready steps
		for s, before := range instructions {
			if len(before) > 0 {
				instructions[s] = Remove(step, instructions[s])

				if len(instructions[s]) == 0 {
					ready = append(ready, s)
					// make sure that the ready steps are ordered alphabetically
					sort.Strings(ready)
				}
			}
		}
	}

	return ordered
}

func main() {
	instr, err := ReadInstructionsFromFile()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(FindOrder(instr))
}
