package queue

import (
	"testing"
)

func TestAddAsec(t *testing.T) {
	q := New()
	q.AddAsec(4)
	q.AddAsec(1)
	q.AddAsec(3)
	q.AddAsec(2)

	expected := []int{1, 2, 3, 4}
	for i, p := 0, q.queue.Front(); i < len(expected); i, p = i+1, p.Next() {
		if p.Value.(int) != expected[i] {
			t.Fatalf("AddAsec failed, expected=%v, actual=%v", expected[i], p.Value.(int))
		}
	}
}

func TestAddDesc(t *testing.T) {
	q := New()
	q.AddDesc(4)
	q.AddDesc(1)
	q.AddDesc(3)
	q.AddDesc(2)

	expected := []int{4, 3, 2, 1}
	for i, p := 0, q.queue.Front(); i < len(expected); i, p = i+1, p.Next() {
		if p.Value.(int) != expected[i] {
			t.Fatalf("AddDesc failed, expected=%v, actual=%v", expected[i], p.Value.(int))
		}
	}
}
