// This is the elevator's low level API
package internal

import (
	"log"
	"sync"
	"time"
)

const (
	StatusIdle    = 0
	StatusRunning = 1
)

var currentFloor = 1
var status = StatusIdle
var mu sync.Mutex

func GetCurrentFloor() int {
	return currentFloor
}

func GetCurrentStatus() int {
	mu.Lock()
	defer mu.Unlock()
	return status
}

func Up(f func()) {
	mu.Lock()
	defer mu.Unlock()
	time.Sleep(1 * time.Second)
	log.Printf("move up, cur=%v", currentFloor)
	currentFloor++
	f()
}

func Down(f func()) {
	mu.Lock()
	defer mu.Unlock()
	time.Sleep(1 * time.Second)
	log.Printf("move down, cur=%v", currentFloor)
	currentFloor--
	f()
}
