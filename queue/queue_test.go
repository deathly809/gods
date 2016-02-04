package queue

import "testing"

func TestQueueNew(t *testing.T) {
	q := New()
	if q.Count() != 0 {
		t.Error("Expected 0 size Queue but found ", q.Count())
	}
}

func TestQueueEnqueue(t *testing.T) {
	q := New()
	for i := 0; i < 100; i++ {
		q.Enqueue(i)
	}
	if q.Count() != 100 {
		t.Error("Expected 100 sized Queue but found ", q.Count())
	}
}

func TestQueueFillThenEmpty(t *testing.T) {
	q := New()
	for i := 0; i < 100; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 100; i++ {
		val := q.Dequeue()
		if val != i {
			t.Errorf("Expected dequeue value of %d but found %d", i,val) 
		}
	}
	if q.Count() != 0 {
		t.Error("Expected empty queue")
	}
}
