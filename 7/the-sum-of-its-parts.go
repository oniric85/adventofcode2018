package main

import (
	"fmt"
	"log"
	"sort"
)

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
