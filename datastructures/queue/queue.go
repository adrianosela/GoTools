package queue

// Queue represents the generic queue
type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
}
