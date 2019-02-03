package queue

import "github.com/adrianosela/GoTools/datastructures/linkedlist"

// LLQueue is the linked-list implementation of a queue
type LLQueue linkedlist.LList

// NewLLQueue starts a new queue with the data of the first node
func NewLLQueue(data interface{}) *LLQueue {
	this := &LLQueue{
		Head: linkedlist.NewNode(data),
	}
	this.Tail = this.Head
	return this
}

// Enqueue puts an item at the back of the queue
// Constant runtime O(1)
func (q *LLQueue) Enqueue(data interface{}) {
	q.Tail.Next = &linkedlist.Node{
		Data: data,
	}
	q.Tail = q.Tail.Next
}

// Dequeue removes the item at the front of the queue
// Constant runtime O(1)
func (q *LLQueue) Dequeue() interface{} {
	rmed := q.Head.Data
	q.Head = q.Head.Next
	// in other languages you might have to delete the node associated
	// with the old "head" node, however golang has a garbage collector
	// which wipes pointer-less memory from the heap
	return rmed
}
