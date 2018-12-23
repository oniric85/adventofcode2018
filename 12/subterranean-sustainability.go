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
	shift = initialShift
	// add padding to simplify logic
	state = ".." + state + ".."
	for i := 2; i < len(state)-2; i++ {
		pattern := state[i-2 : i+3]
		outcome := rules.Match(pattern)
		newState += outcome
	}

	if newState[0:1] == "#" {
		newState = "." + newState
		shift--
	}

	if newState[len(newState)-1:] == "#" {
		newState = newState + "."
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

func main() {
	state, ruleset, err := ReadInstructions()

	if err != nil {
		log.Fatal(err)
	}

	shift := 0
	fmt.Println(" 0:", state)
	for i := 1; i <= 20; i++ {
		state, shift = GrowGeneration(state, ruleset, shift)
		fmt.Printf("%2d: %s\n", i, state)
	}

	fmt.Println("Result:", Result(state, shift))
}
