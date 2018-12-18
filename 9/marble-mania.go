package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

type MarblesGame struct {
	totalMarbles  int
	marbles       []int
	players       []int
	highScore     int
	currentMarble int
	currentValue  int
	currentPlayer int
}

func Reminder(num int, den int) int {
	r := num % den
	if r < 0 {
		return r + den
	}

	return r
}

func (game *MarblesGame) New(numPlayers int, totalMarbles int) {
	game.players = make([]int, numPlayers)
	game.totalMarbles = totalMarbles

	game.marbles = []int{0}

	game.currentMarble = 0
	game.currentValue = 1
	game.currentPlayer = 0
}

func (game *MarblesGame) PositionCurrentMarble() {
	nextPosition := ((game.currentMarble+1)%len(game.marbles) + 1) % (len(game.marbles) + 1)

	game.marbles = append(game.marbles, 0) // make room for a new element
	copy(game.marbles[(nextPosition+1)%len(game.marbles):], game.marbles[nextPosition:])
	game.marbles[nextPosition] = game.currentValue

	game.currentMarble = nextPosition
}

func (game *MarblesGame) ScorePoints() {
	toBeRemovedMarble := Reminder(game.currentMarble-7, len(game.marbles))

	game.players[game.currentPlayer] += game.currentValue
	game.players[game.currentPlayer] += game.marbles[toBeRemovedMarble]

	if game.players[game.currentPlayer] > game.highScore {
		game.highScore = game.players[game.currentPlayer]
	}

	game.marbles = append(game.marbles[:toBeRemovedMarble], game.marbles[toBeRemovedMarble+1:]...)
	game.currentMarble = toBeRemovedMarble
}

func (game *MarblesGame) Move() {
	if game.currentValue%23 == 0 {
		game.ScorePoints()
	} else {
		game.PositionCurrentMarble()
	}

	game.currentValue++
	game.currentPlayer = (game.currentPlayer + 1) % len(game.players)
}

func (game *MarblesGame) GetHighScore() int {
	return game.highScore
}

func (game *MarblesGame) Play() {
	for i := 1; i <= game.totalMarbles; i++ {
		game.Move()
	}
}

func ReadParametersFromFile() (numPlayers int, lastMarbleValue int, err error) {
	bytes, err := ioutil.ReadFile("input.txt")

	if err != nil {
		return numPlayers, lastMarbleValue, err
	}

	s := string(bytes)

	r, _ := regexp.Compile("^([0-9]+) players; last marble is worth ([0-9]+) points$")

	matches := r.FindStringSubmatch(s)

	numPlayers, _ = strconv.Atoi(matches[1])
	lastMarbleValue, _ = strconv.Atoi(matches[2])

	return numPlayers, lastMarbleValue, nil
}

func main() {
	numPlayers, lastMarbleValue, err := ReadParametersFromFile()

	if err != nil {
		log.Fatal(err)
	}

	game := MarblesGame{}
	game.New(numPlayers, lastMarbleValue)
	game.Play()

	fmt.Println("High score is:", game.GetHighScore())

	gameSecondPart := MarblesGame{}
	game.New(numPlayers, lastMarbleValue*100)
	game.Play()
	fmt.Println("High score for second part of the puzzle is:", gameSecondPart.GetHighScore())
}
