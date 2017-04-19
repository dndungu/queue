package queue

type node struct {
	next  *node
	value interface{}
}

// Queue holds a pointer to the head and the tail of the queue data structure
type Queue struct {
	head   *node
	length int64
	tail   *node
}

// NewQueue creates an new instance of a queue data structure
func NewQueue() *Queue {
	return &Queue{}
}

// Empty returns true when the Queue does not contain any nodes and false when it does
func (q *Queue) Empty() bool {
	return q.length == 0
}

// Dequeue removes and returns the node at the tail
func (q *Queue) Dequeue() (interface{}, bool) {
	if q.Empty() {
		return nil, false
	}
	defer func() { q.length-- }()
	v := q.tail.value
	if q.tail.next == nil {
		q.tail = nil
		q.head = nil
	} else {
		q.tail = q.tail.next
	}
	return v, true
}

// Peek returns the value of the node at the tail
func (q *Queue) Peek() (interface{}, bool) {
	if q.Empty() {
		return nil, false
	}
	return q.tail.value, true
}

// Enqueue adds a new node to the head of the Queue
func (q *Queue) Enqueue(v interface{}) {
	defer func() { q.length++ }()
	node := &node{nil, v}
	if q.Empty() {
		q.head = node
		q.tail = node
		return
	}
	q.head.next = node
	q.head = node
}

// Length returns the number of elements in the Queue
func (q *Queue) Length() int64 {
	return q.length
}
