package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

type Marble struct {
	value uint
	prev  *Marble
	next  *Marble
}

type MarblesGame struct {
	totalMarbles  int
	players       []uint
	currentMarble *Marble
	currentValue  uint
	currentPlayer int
	highScore     uint
}

func (game *MarblesGame) New(numPlayers int, totalMarbles int) {
	game.players = make([]uint, numPlayers)
	game.totalMarbles = totalMarbles

	game.currentMarble = &Marble{value: 0}
	game.currentMarble.prev, game.currentMarble.next = game.currentMarble, game.currentMarble
	game.currentValue = 1
	game.currentPlayer = 0
}

func (game *MarblesGame) PositionNewMarble() {
	newMarble := &Marble{value: game.currentValue}
	newMarble.prev = game.currentMarble.next
	newMarble.next = game.currentMarble.next.next

	game.currentMarble.next.next.prev = newMarble
	game.currentMarble.next.next = newMarble

	game.currentMarble = newMarble
}

func (game *MarblesGame) MoveCurrentCounterClockwise(num int) {
	for i := 0; i < num; i++ {
		game.currentMarble = game.currentMarble.prev
	}
}

func (game *MarblesGame) RemoveCurrent() {
	current := game.currentMarble.next
	current.prev = game.currentMarble.prev
	current.prev.next = current

	game.currentMarble = current
}

func (game *MarblesGame) ScorePoints() {
	game.players[game.currentPlayer] += game.currentValue
	game.MoveCurrentCounterClockwise(7)
	game.players[game.currentPlayer] += game.currentMarble.value

	if game.players[game.currentPlayer] > game.highScore {
		game.highScore = game.players[game.currentPlayer]
	}

	game.RemoveCurrent()
}

func (game *MarblesGame) Move() {
	if game.currentValue%23 == 0 {
		game.ScorePoints()
	} else {
		game.PositionNewMarble()
	}

	game.currentValue++
	game.currentPlayer = (game.currentPlayer + 1) % len(game.players)
}

func (game *MarblesGame) GetHighScore() uint {
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
	gameSecondPart.New(numPlayers, lastMarbleValue*100)
	gameSecondPart.Play()
	fmt.Println("High score for second part of the puzzle is:", gameSecondPart.GetHighScore())
}
