
package queue

// Queue gives required methods for a queue data-structure
type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	Count() int
} 

type queue []interface{}

func (q *queue) Enqueue(val interface{}) {
	*q = append(*q,val)
}

func (q *queue) Dequeue() interface{} {
	if len(*q) > 0{
		v := (*q)[0]
		*q = (*q)[1:]
		return v
	}
	return nil
}

func (q *queue) Count() int {
	return len(*q)
}

// New constructs a new queue
func New() Queue {
	return new(queue)
}