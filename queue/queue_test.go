package queue

import (
	"testing"
)

func TestAddAsc1(t *testing.T) {
	q := New()
	q.AddAsc(&Task{4, 1}, nil)
	q.AddAsc(&Task{1, 1}, nil)
	q.AddAsc(&Task{3, 1}, nil)
	q.AddAsc(&Task{2, 1}, nil)

	expected := []int{1, 2, 3, 4}
	for i, p := 0, q.queue.Front(); i < len(expected); i, p = i+1, p.Next() {
		if p.Value.(*Task).TargetFloor != expected[i] {
			t.Fatalf("AddAsc failed, expected=%v, actual=%v", expected[i], p.Value.(*Task).TargetFloor)
		}
	}
}

func TestAddAsc2(t *testing.T) {
	q := New()
	q.AddAsc(&Task{5, 2}, nil)
	q.AddAsc(&Task{10, 1}, nil)

	expected := []int{10}
	for i, p := 0, q.queue.Front(); i < len(expected); i, p = i+1, p.Next() {
		if p.Value.(*Task).TargetFloor != expected[i] {
			t.Fatalf("AddAsc failed, expected=%v, actual=%v", expected[i], p.Value.(*Task).TargetFloor)
		}
	}
}

func TestAddAsc3(t *testing.T) {
	q := New()
	q.AddAsc(&Task{5, 2}, nil)
	q.AddAsc(&Task{10, 1}, nil)
	q.AddAsc(&Task{6, 1}, nil)
	q.AddAsc(&Task{8, 2}, nil)
	q.AddAsc(&Task{12, 1}, nil)

	expected := []int{6, 10, 12}
	for i, p := 0, q.queue.Front(); i < len(expected); i, p = i+1, p.Next() {
		if p.Value.(*Task).TargetFloor != expected[i] {
			t.Errorf("AddAsc failed, expected=%v, actual=%v", expected[i], p.Value.(*Task).TargetFloor)
		}
	}
}

func TestAddAsc4(t *testing.T) {
	q := New()
	q.AddAsc(&Task{10, 2}, nil)
	q.AddAsc(&Task{5, 2}, nil)

	expected := []int{10}
	//for p, i := q.queue.Front(), 0; p != nil; p, i = p.Next(), i+1 {
	//	log.Printf("print i=%v, value=%v, direction=%v", i, p.Value.(*Task).TargetFloor, p.Value.(*Task).Direction)
	//}
	for i, p := 0, q.queue.Front(); i < len(expected); i, p = i+1, p.Next() {
		if p.Value.(*Task).TargetFloor != expected[i] {
			t.Fatalf("AddDesc failed, expected=%v, actual=%v", expected[i], p.Value.(*Task).TargetFloor)
		}
	}
}

func TestAddAsc5(t *testing.T) {
	q := New()
	q.AddAsc(&Task{10, 2}, nil)
	q.AddAsc(&Task{5, 2}, nil)
	q.AddAsc(&Task{4, 1}, nil)

	expected := []int{4, 10}
	//for p, i := q.queue.Front(), 0; p != nil; p, i = p.Next(), i+1 {
	//	log.Printf("print i=%v, value=%v, direction=%v", i, p.Value.(*Task).TargetFloor, p.Value.(*Task).Direction)
	//}
	for i, p := 0, q.queue.Front(); i < len(expected); i, p = i+1, p.Next() {
		if p.Value.(*Task).TargetFloor != expected[i] {
			t.Fatalf("AddDesc failed, expected=%v, actual=%v", expected[i], p.Value.(*Task).TargetFloor)
		}
	}
}

func TestAddDesc1(t *testing.T) {
	q := New()
	q.AddDesc(&Task{4, 2}, nil)
	q.AddDesc(&Task{1, 2}, nil)
	q.AddDesc(&Task{3, 2}, nil)
	q.AddDesc(&Task{2, 2}, nil)

	expected := []int{4, 3, 2, 1}
	for i, p := 0, q.queue.Front(); i < len(expected); i, p = i+1, p.Next() {
		if p.Value.(*Task).TargetFloor != expected[i] {
			t.Fatalf("AddDesc failed, expected=%v, actual=%v", expected[i], p.Value.(*Task).TargetFloor)
		}
	}
}

func TestAddDesc2(t *testing.T) {
	q := New()
	q.AddDesc(&Task{5, 1}, nil)
	q.AddDesc(&Task{10, 2}, nil)

	expected := []int{10}
	for i, p := 0, q.queue.Front(); i < len(expected); i, p = i+1, p.Next() {
		if p.Value.(*Task).TargetFloor != expected[i] {
			t.Fatalf("AddDesc failed, expected=%v, actual=%v", expected[i], p.Value.(*Task).TargetFloor)
		}
	}
}

func TestAddDesc3(t *testing.T) {
	q := New()
	q.AddDesc(&Task{5, 1}, nil)
	q.AddDesc(&Task{10, 2}, nil)
	q.AddDesc(&Task{6, 2}, nil)
	q.AddDesc(&Task{8, 1}, nil)
	q.AddDesc(&Task{12, 2}, nil)

	expected := []int{12, 10, 6}
	for i, p := 0, q.queue.Front(); i < len(expected); i, p = i+1, p.Next() {
		if p.Value.(*Task).TargetFloor != expected[i] {
			t.Errorf("AddDesc failed, expected=%v, actual=%v", expected[i], p.Value.(*Task).TargetFloor)
		}
	}
}

func TestAddDesc4(t *testing.T) {
	q := New()
	q.AddDesc(&Task{10, 1}, nil)
	q.AddDesc(&Task{5, 1}, nil)

	expected := []int{5}
	//for p, i := q.queue.Front(), 0; p != nil; p, i = p.Next(), i+1 {
	//	log.Printf("print i=%v, value=%v, direction=%v", i, p.Value.(*Task).TargetFloor, p.Value.(*Task).Direction)
	//}
	for i, p := 0, q.queue.Front(); i < len(expected); i, p = i+1, p.Next() {
		if p.Value.(*Task).TargetFloor != expected[i] {
			t.Fatalf("AddDesc failed, expected=%v, actual=%v", expected[i], p.Value.(*Task).TargetFloor)
		}
	}
}

func TestAddDesc5(t *testing.T) {
	q := New()
	q.AddDesc(&Task{10, 1}, nil)
	q.AddDesc(&Task{5, 1}, nil)
	q.AddDesc(&Task{4, 2}, nil)

	expected := []int{4}
	//for p, i := q.queue.Front(), 0; p != nil; p, i = p.Next(), i+1 {
	//	log.Printf("print i=%v, value=%v, direction=%v", i, p.Value.(*Task).TargetFloor, p.Value.(*Task).Direction)
	//}
	for i, p := 0, q.queue.Front(); i < len(expected); i, p = i+1, p.Next() {
		if p.Value.(*Task).TargetFloor != expected[i] {
			t.Fatalf("AddDesc failed, expected=%v, actual=%v", expected[i], p.Value.(*Task).TargetFloor)
		}
	}
}
