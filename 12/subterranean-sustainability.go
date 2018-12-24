package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Rule struct {
	pattern string
	outcome string
}

type RuleSet struct {
	rules []Rule
}

func (r *RuleSet) Append(rule Rule) {
	r.rules = append(r.rules, rule)
}

func (r *RuleSet) Match(str string) string {
	for _, rule := range r.rules {
		if str == rule.pattern {
			return rule.outcome
		}
	}

	// default is no plant
	return "."
}

func GrowGeneration(state string, rules RuleSet, initialShift int) (newState string, shift int) {
	shift = initialShift - 3
	// add padding to simplify logic
	state = "....." + state + "....."
	for i := 2; i < len(state)-2; i++ {
		pattern := state[i-2 : i+3]
		outcome := rules.Match(pattern)
		newState += outcome
	}

	// adjust borders and shift
	for newState[0] == '.' {
		shift++
		newState = newState[1:]
	}

	for newState[len(newState)-1] == '.' {
		newState = newState[:len(newState)-1]
	}

	return newState, shift
}

func Result(state string, shift int) (result int) {
	for pos, c := range state {
		if string(c) == "#" {
			result += pos + shift
		}
	}
	return result
}

func ReadInstructions() (string, RuleSet, error) {
	file, err := os.Open("input.txt")

	ruleset := RuleSet{}

	if err != nil {
		return "", ruleset, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var state string
	r, _ := regexp.Compile("^([\\.#]{5}) => (#|\\.)$")
	for scanner.Scan() {
		if state == "" {
			state = scanner.Text()[len("initial state: "):]
		} else {
			matches := r.FindStringSubmatch(scanner.Text())
			if len(matches) > 0 {
				rule := Rule{pattern: matches[1], outcome: matches[2]}
				ruleset.Append(rule)
			}
		}
	}

	return state, ruleset, scanner.Err()
}

func Answer(state string, ruleset RuleSet, generations int) int {
	shift, prevShift := 0, 0
	prevResult := 0
	states := make(map[string]bool)
	for i := 1; i <= generations; i++ {
		state, shift = GrowGeneration(state, ruleset, prevShift)
		if _, ok := states[state]; ok {
			// calculate final result based on linear regression
			result := Result(state, shift)
			return result + (generations-i)*(result-prevResult)
		}
		states[state] = true
		prevShift = shift
		prevResult = Result(state, shift)
	}

	return Result(state, shift)
}

func main() {
	state, ruleset, err := ReadInstructions()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result for 20 generations is:", Answer(state, ruleset, 20))

	// my input start looping at generation 186
	// after the loop the pattern kept shifting right one position every generation
	// and the result increased by 194 on each iteration
	fmt.Println("Result for 50000000000 generations is:", Answer(state, ruleset, 50000000000))
}
