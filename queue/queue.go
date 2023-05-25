// This package implements a queue data structure which is used for task queue.
package queue

import (
	"container/list"

	"log"
)

type Queue struct {
	queue *list.List
}

func New() *Queue {
	var q Queue
	q.queue = list.New()
	return &q
}

// Add an element into the queue in asec order
func (q *Queue) AddAsec(v int) {
	// empty list
	if q.queue.Len() == 0 {
		q.queue.PushBack(v)
		return
	}

	// insert the v into current position by asec
	var p *list.Element
	for p = q.queue.Front(); p.Next() != nil; p = p.Next() {
		curValue := p.Value.(int)
		// don't insert duplicate value
		if v == curValue {
			return
		} else if v < curValue {
			q.queue.InsertBefore(v, p)
			return
		}
	}
	if p.Next() == nil {
		if v > p.Value.(int) {
			q.queue.InsertAfter(v, p)
		} else if v < p.Value.(int) {
			q.queue.InsertBefore(v, p)
		}
	}
}

// Add an element into the queue in desc order
func (q *Queue) AddDesc(v int) {
	// empty list
	if q.queue.Len() == 0 {
		q.queue.PushBack(v)
		return
	}

	// insert the v into current position by asec
	var p *list.Element
	for p = q.queue.Front(); p.Next() != nil; p = p.Next() {
		curValue := p.Value.(int)
		// don't insert duplicate value
		if v == curValue {
			return
		} else if v > curValue {
			q.queue.InsertBefore(v, p)
		}
	}
	if p.Next() == nil {
		if v > p.Value.(int) {
			q.queue.InsertBefore(v, p)
		} else if v < p.Value.(int) {
			q.queue.InsertAfter(v, p)
		}
	}
}

// Get the first element's value
func (q *Queue) Front() int {
	return q.queue.Front().Value.(int)
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
