// This package implements a queue data structure which is used for task queue.
package queue

import (
	"container/list"

	"log"
)

type Queue struct {
	queue *list.List
}

const (
	Up   = 1
	Down = 2
)

type Task struct {
	TargetFloor int
	Direction   int
}

func New() *Queue {
	var q Queue
	q.queue = list.New()
	return &q
}

// Add an element into the queue in asec order
func (q *Queue) AddAsc(v *Task, downQ *Queue) []*Task {
	// empty list
	if q.queue.Len() == 0 {
		q.queue.PushBack(v)
		return nil
	}

	// Tasks to be removed from upTaskQueue to downTaskQueue
	var removedTasks []*Task

	// insert the v into current position by asec
	var p *list.Element
	for p = q.queue.Front(); p.Next() != nil; {
		curValue := p.Value.(*Task).TargetFloor
		curDirection := p.Value.(*Task).Direction
		// don't insert duplicate value
		if v.TargetFloor == curValue {
			if v.Direction == curDirection { // already exists, don't add duplicate tasks
				return removedTasks
			} else if v.Direction == Up { // replace the down task with up task
				p.Value.(*Task).Direction = Up
				removedTasks = append(removedTasks, &Task{p.Value.(*Task).TargetFloor, p.Value.(*Task).Direction})
			}
			p = p.Next()
		} else if v.TargetFloor < curValue {
			if v.Direction == Up {
				q.queue.InsertBefore(v, p)
			} else {
				removedTasks = append(removedTasks, v)
			}
			return removedTasks
		} else {
			if p.Value.(*Task).Direction == Down {
				removedTasks = append(removedTasks, &Task{p.Value.(*Task).TargetFloor, p.Value.(*Task).Direction})
				next := p.Next()
				q.queue.Remove(p)
				p = next
			} else {
				p = p.Next()
			}
		}
	}

	// handle the last node
	if p.Next() == nil {
		if v.TargetFloor > p.Value.(*Task).TargetFloor {
			q.queue.InsertAfter(v, p)
			if p.Value.(*Task).Direction == Down {
				removedTasks = append(removedTasks, &Task{p.Value.(*Task).TargetFloor, p.Value.(*Task).Direction})
				q.queue.Remove(p)
			}
		} else if v.TargetFloor < p.Value.(*Task).TargetFloor {
			if v.Direction == Up {
				q.queue.InsertBefore(v, p)
			} else {
				removedTasks = append(removedTasks, v)
			}
		}
	}

	if downQ != nil {
		for _, v := range removedTasks {
			downQ.AddDesc(v, q)
		}
	}

	return removedTasks
}

// Add an element into the queue in asec order
func (q *Queue) AddDesc(v *Task, upQ *Queue) []*Task {
	// empty list
	if q.queue.Len() == 0 {
		q.queue.PushBack(v)
		return nil
	}

	// Tasks to be removed from upTaskQueue to downTaskQueue
	var removedTasks []*Task

	// insert the v into current position by asec
	var p *list.Element
	for p = q.queue.Front(); p.Next() != nil; {
		curValue := p.Value.(*Task).TargetFloor
		curDirection := p.Value.(*Task).Direction
		// don't insert duplicate value
		if v.TargetFloor == curValue {
			if v.Direction == curDirection { // already exists, don't add duplicate tasks
				return removedTasks
			} else if v.Direction == Down { // replace the up task with up task
				p.Value.(*Task).Direction = Down
				removedTasks = append(removedTasks, &Task{p.Value.(*Task).TargetFloor, p.Value.(*Task).Direction})
			}
			p = p.Next()
		} else if v.TargetFloor > curValue {
			if v.Direction == Down {
				q.queue.InsertBefore(v, p)
			} else {
				removedTasks = append(removedTasks, v)
			}
			return removedTasks
		} else {
			if p.Value.(*Task).Direction == Up {
				removedTasks = append(removedTasks, &Task{p.Value.(*Task).TargetFloor, p.Value.(*Task).Direction})
				next := p.Next()
				q.queue.Remove(p)
				p = next
			} else {
				p = p.Next()
			}
		}
	}

	// handle the last node
	if p.Next() == nil {
		if v.TargetFloor < p.Value.(*Task).TargetFloor {
			q.queue.InsertAfter(v, p)
			if p.Value.(*Task).Direction == Up {
				removedTasks = append(removedTasks, &Task{p.Value.(*Task).TargetFloor, p.Value.(*Task).Direction})
				q.queue.Remove(p)
			}
		} else if v.TargetFloor > p.Value.(*Task).TargetFloor {
			if v.Direction == Down {
				q.queue.InsertBefore(v, p)
			} else {
				removedTasks = append(removedTasks, v)
			}
		}
	}

	if upQ != nil {
		for _, v := range removedTasks {
			upQ.AddAsc(v, q)
		}
	}

	return removedTasks
}

// Get the first element's value
func (q *Queue) Front() int {
	return q.queue.Front().Value.(*Task).TargetFloor
}

// Popup the first element
func (q *Queue) PopFront() {
	f := q.queue.Front()
	if f == nil {
		log.Printf("queue shouln't be empty")
		return
	}
	q.queue.Remove(f)
}

// Get the queue's current length
func (q *Queue) Len() int {
	return q.queue.Len()
}
