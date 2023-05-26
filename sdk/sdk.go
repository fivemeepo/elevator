package sdk

import (
	"log"
	"sync"
	"test/elevator/internal"
	"test/elevator/queue"
)

const (
	Stop = 0
	Up   = 1
	Down = 2
)

type ElevatorSDK interface {
	Up(curFloor int)      // button "up"
	Down(curFloor int)    // button "down"
	Goto(targetFloor int) // button "number"
}

type Elevator struct {
	mu           sync.Mutex
	curDirection int           // 0:stop, 1: up, 2:down
	upTasks      *queue.Queue  // upward task queue
	downTasks    *queue.Queue  // downward task queue
	signal       chan struct{} // elevator starts to run when receiving signal
}

// Must call Init() in main()
func (e *Elevator) Init() {
	e.upTasks = queue.New()
	e.downTasks = queue.New()
	go e.schedule()
	go e.run()
}

// Up button's handler
func (e *Elevator) Up(curFloor int) {
	// return immediatelly and start a new goroutine to handle the request asynchronously
	log.Printf("receive up req, floor=%v", curFloor)
	go e.addTask(curFloor, Up)
}

// Down button's handler
func (e *Elevator) Down(curFloor int) {
	// return immediatelly and start a new goroutine to handle the request asynchronously
	log.Printf("receive down req, floor=%v", curFloor)
	go e.addTask(curFloor, Down)
}

// Goto should only be pressed after the person is already picked up.
func (e *Elevator) Goto(curFloor int) {
	log.Printf("receive goto req, floor=%v", curFloor)
	if curFloor > internal.GetCurrentFloor() {
		go e.addTask(curFloor, Up)
	} else {
		go e.addTask(curFloor, Down)
	}
}

// When a user presses a button, add a task to the elevator's task queue
func (e *Elevator) addTask(curFloor int, button int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	EleFloor := internal.GetCurrentFloor()
	var targetDirection int

	// check current direction, and add new task to up/down task queue
	switch e.curDirection {
	case Stop:
		if curFloor > EleFloor {
			targetDirection = Up
			e.upTasks.AddAsc(&queue.Task{curFloor, button}, e.downTasks)
		} else if curFloor < EleFloor {
			targetDirection = Down
			e.downTasks.AddDesc(&queue.Task{curFloor, button}, e.upTasks)
		}
	case Up:
		if curFloor > EleFloor {
			if button == Up {
				e.upTasks.AddAsc(&queue.Task{curFloor, button}, e.downTasks)
			} else {
				e.downTasks.AddDesc(&queue.Task{curFloor, button}, e.upTasks)
			}
		} else {
			e.downTasks.AddDesc(&queue.Task{curFloor, button}, e.upTasks)
		}
	case Down:
		if curFloor < EleFloor {
			if button == Down {
				e.downTasks.AddDesc(&queue.Task{curFloor, button}, e.upTasks)
			} else {
				e.upTasks.AddAsc(&queue.Task{curFloor, button}, e.downTasks)
			}
		} else {
			e.upTasks.AddAsc(&queue.Task{curFloor, button}, e.downTasks)
		}
	default:
		log.Fatal("unexpected direction")
	}

	// if status changes, then call schedule to start the elevator
	if e.curDirection == Stop {
		e.curDirection = targetDirection
		e.schedule()
	}
}

// Trigger the elevator to run
// it will be used as the callback function for internal.Up()/Down()
func (e *Elevator) schedule() {
	// lazy init signal
	if e.signal == nil {
		e.signal = make(chan struct{})
	}

	// send a signal to start the elevator
	e.signal <- struct{}{}
}

// Move once and wait until the next signal comes
func (e *Elevator) run() {
	// lazy init signal
	if e.signal == nil {
		e.signal = make(chan struct{})
	}

	// start to handle tasks
	for {
		//log.Printf("waiting for signal")
		<-e.signal
		e.mu.Lock()

		// check elevator's current direction
		switch e.curDirection {
		case Stop:
			log.Printf("stop at floor=%v, waitint for new requests", internal.GetCurrentFloor())

		// if the elevator is going up, check whehter it arrives at the destination, if yes, remove the task from task queue
		// if no more tasks left, then change the elevator's status to idle(stop)
		case Up:
			curFloor := internal.GetCurrentFloor()
			targetFloor := e.upTasks.Front()

			if curFloor < targetFloor { // move up
				go internal.Up(e.schedule)
			} else if curFloor == targetFloor { // arrive at a target floor, check whether to continue or not
				e.upTasks.PopFront()
				if e.upTasks.Len() == 0 && e.downTasks.Len() == 0 { // no more tasks, stop
					e.curDirection = Stop
					log.Printf("go up to floor=%v, stop", internal.GetCurrentFloor())
				} else if e.upTasks.Len() == 0 && e.downTasks.Len() != 0 { // turn around
					e.curDirection = Down
					log.Printf("turn around to go down from floor=%v", internal.GetCurrentFloor())
					go e.schedule()
				} else { // keep going up
					log.Printf("go up to floor=%v, stop and continue", internal.GetCurrentFloor())
					go e.schedule()
				}
			} else {
				log.Printf("target floor=%v shouldn't be lower than current floor=%v when going up", curFloor, targetFloor)
				e.mu.Unlock()
				panic(0)
			}

		// same as going up
		case Down:
			curFloor := internal.GetCurrentFloor()
			targetFloor := e.downTasks.Front()
			if curFloor > targetFloor {
				go internal.Down(e.schedule)
			} else if curFloor == targetFloor {
				e.downTasks.PopFront()
				if e.upTasks.Len() == 0 && e.downTasks.Len() == 0 {
					e.curDirection = Stop
					log.Printf("go down to floor=%v, stop", internal.GetCurrentFloor())
				} else if e.downTasks.Len() == 0 && e.upTasks.Len() != 0 {
					e.curDirection = Up
					log.Printf("turn around to go up from floor=%v", internal.GetCurrentFloor())
					go e.schedule()
				} else {
					log.Printf("go down to floor=%v, stop and continue", internal.GetCurrentFloor())
					go e.schedule()
				}
			} else {
				log.Printf("target floor=%v shouldn't be lower than current floor=%v when going up", curFloor, targetFloor)
				e.mu.Unlock()
				panic(0)
			}
		}
		e.mu.Unlock()
	}

}
