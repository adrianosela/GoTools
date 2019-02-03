package queue

import "github.com/adrianosela/GoTools/datastructures/linkedlist"

// LLQueue is the linked-list implementation of a queue
type LLQueue struct {
	Front *linkedlist.Node
	Back  *linkedlist.Node
}

// NewLLQueue starts a new queue with the data of the first node
func NewLLQueue(data interface{}) *LLQueue {
	this := &LLQueue{
		Front: linkedlist.NewNode(data),
	}
	this.Back = this.Front
	return this
}

// Enqueue puts an item at the back of the queue
// Constant runtime O(1)
func (q *LLQueue) Enqueue(data interface{}) {
	q.Back.Next = &linkedlist.Node{
		Data: data,
	}
	q.Back = q.Back.Next
}

// Dequeue removes the item at the front of the queue
// Constant runtime O(1)
func (q *LLQueue) Dequeue() interface{} {
	rmed := q.Front.Data
	q.Front = q.Front.Next
	// in other languages you might have to delete the node associated
	// with the old "front" node, however golang has a garbage collector
	// which wipes pointer-less memory from the heap
	return rmed
}
