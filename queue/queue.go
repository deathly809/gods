package queue

type queueNode struct {
	prev, next *queueNode
	data       interface{}
}

// Queue gives required methods for a queue data-structure
type Queue interface {
	Enqueue(interface{})
	Peek() interface{}
	Dequeue() interface{}
	Count() int
}

type queue struct {
	head, tail *queueNode
	count      int
}

//type queue []interface{}

func (q *queue) Enqueue(val interface{}) {
	if q.count == 0 {
		newNode := &queueNode{}
		newNode.data = val
		q.head, q.tail = newNode, newNode
	} else {
		newNode := &queueNode{prev: q.tail, data: val}
		q.tail.next = newNode
		q.tail = newNode
	}
	q.count++
}

func (q *queue) Dequeue() interface{} {
	result := interface{}(nil)
	if q.head != nil {

		result = q.head.data
		q.head = q.head.next

		if q.head == nil {
			q.tail = nil
		} else {
			q.head.prev = nil
		}
		q.count--
	}
	return result
}

func (q *queue) Peek() interface{} {
	result := interface{}(nil)
	if q.head != nil {
		result = q.head.data
	}
	return result
}

func (q *queue) Count() int {
	return q.count
}

// New constructs a new queue
func New() Queue {
	return new(queue)
}
