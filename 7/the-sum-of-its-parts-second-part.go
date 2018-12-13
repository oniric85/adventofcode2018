package main

import (
	"fmt"
	"log"
	"sort"
)

type Worker struct {
	step string
	work int
}

func (worker *Worker) IsIdle() bool {
	return worker.work == 0
}

func (worker *Worker) Work() {
	worker.work--
}

func (worker *Worker) Remaining() int {
	return worker.work
}

func (worker *Worker) Assign(step string) {
	worker.step = step
	worker.work = 60 + int(byte(step[0])) - 65 + 1 // 65 = A in ASCII
}

type WorkQueue struct {
	workers []*Worker
}

func (q *WorkQueue) Len() int {
	return len(q.workers)
}

func (q *WorkQueue) Get(pos int) *Worker {
	return q.workers[pos]
}

func (q *WorkQueue) Remove(pos int) *Worker {
	w := q.workers[pos]
	q.workers = append(q.workers[:pos], q.workers[pos+1:]...)
	return w
}

func (q *WorkQueue) New(initial int) *WorkQueue {
	q.workers = make([]*Worker, initial)

	for i := 0; i < initial; i++ {
		q.workers[i] = new(Worker)
	}

	return q
}

func (q *WorkQueue) Enqueue(worker *Worker) {
	q.workers = append(q.workers, worker)
}

func (q *WorkQueue) Dequeue() (worker *Worker) {
	w := q.workers[0]
	q.workers = q.workers[1:]
	return w
}

func (q *WorkQueue) IsEmpty() bool {
	return len(q.workers) == 0
}

func FindIdleWorker(workers []*Worker) *Worker {
	for _, worker := range workers {
		if worker.IsIdle() {
			return worker
		}
	}

	return nil
}

func FindTotalTimeWithHelpFromElves(instructions map[string][]string, elves int) int {
	ready := FindReadySteps(instructions)

	var step string

	idleWorkers := WorkQueue{}
	idleWorkers.New(elves)

	workingWorkers := WorkQueue{}
	workingWorkers.New(0) // initially everyone is at rest

	timeElapsed := 0

	ordered := ""
	totalSteps := len(instructions)

	for len(ordered) < totalSteps {
		// check if there is assignable work!
		for !idleWorkers.IsEmpty() && len(ready) > 0 {
			worker := idleWorkers.Dequeue()
			step, ready = Shift(ready)
			worker.Assign(step)
			workingWorkers.Enqueue(worker)
		}

		fmt.Print(timeElapsed, " ")

		for i := 0; i < workingWorkers.Len(); i++ {
			worker := workingWorkers.Get(i)

			fmt.Print(worker.step, "(", worker.Remaining(), ") ")
		}

		fmt.Println("|", ordered)

		for i := 0; i < workingWorkers.Len(); i++ {
			worker := workingWorkers.Get(i)
			worker.Work()

			if worker.IsIdle() {
				workingWorkers.Remove(i)
				i-- // we removed an element
				idleWorkers.Enqueue(worker)
				step := worker.step
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
		}

		timeElapsed++
	}

	return timeElapsed
}

func main() {
	instr, err := ReadInstructionsFromFile()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total elapsed time is:", FindTotalTimeWithHelpFromElves(instr, 5))
}
