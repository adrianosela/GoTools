package queue

// ArrQueue is the array implementation of a queue
type ArrQueue struct {
	Q     []interface{}
	Front int
	Back  int
	Size  int
}

const defaultInitSize = 10

// NewArrQueue starts a new queue with the data of the first node
func NewArrQueue(data interface{}) *ArrQueue {
	return &ArrQueue{
		Q:     make([]interface{}, defaultInitSize),
		Front: 0,
		Back:  0,
		Size:  defaultInitSize,
	}
}

// Enqueue puts an item at the back of the queue
// Constant runtime O(1) - (amortized)
func (q *ArrQueue) Enqueue(data interface{}) {
	if q.isFull() {
		// dont be fooled, resizing takes time!
		q.Q = append(q.Q, data)
		q.Size++
	} else {
		q.Q[q.Back] = data
	}
	q.Back = (q.Back + 1) % q.Size
}

// Dequeue removes the item at the front of the queue
// Constant runtime O(1)
func (q *ArrQueue) Dequeue() interface{} {
	rmed := q.Q[q.Front]
	q.Front = (q.Front + 1) % q.Size
	return rmed
}

func (q *ArrQueue) isFull() bool {
	return (q.Front == (q.Back+1)%q.Size)
}
