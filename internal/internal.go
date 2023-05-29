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
var sleepTime = 1 * time.Second
var history []int

func GetCurrentFloor() int {
	mu.Lock()
	defer mu.Unlock()
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
	time.Sleep(sleepTime)
	currentFloor++
	log.Printf("move up, cur=%v", currentFloor)
	history = append(history, currentFloor)
	f()
}

func Down(f func()) {
	mu.Lock()
	defer mu.Unlock()
	time.Sleep(sleepTime)
	currentFloor--
	log.Printf("move down, cur=%v", currentFloor)
	history = append(history, currentFloor)
	f()
}

func ShowHistory() {
	log.Printf("History:%v", history)
}
