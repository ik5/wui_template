package tasks

import "container/heap"

// SimpleTask holds details about task to do
type SimpleTask struct {
	DBID     int64
	Value    string
	Priority int
	index    int
}

// SimplePriorityQueue holds a list of tasks
type SimplePriorityQueue []*SimpleTask

// Len is the length of a simple priority queue (interface function)
func (spq SimplePriorityQueue) Len() int {
	return len(spq)
}

// Less helps Pop to figure out  what is the highest priority (interface function)
func (spq SimplePriorityQueue) Less(i, j int) bool {
	return spq[i].Priority > spq[j].Priority
}

// Swap changes the position of two jobs (interface function)
func (spq SimplePriorityQueue) Swap(i, j int) {
	spq[i], spq[j] = spq[j], spq[i]
	spq[i].index = i
	spq[j].index = j
}

// Push pushes a new item to the end of the queue (interface function)
func (spq *SimplePriorityQueue) Push(x interface{}) {
	n := len(*spq)
	item := x.(*SimpleTask)
	item.index = n
	*spq = append(*spq, item)
}

// Pop the last item at the queue (interface function)
func (spq *SimplePriorityQueue) Pop() interface{} {
	old := *spq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*spq = old[0 : n-1]
	return item
}

// Update the current task
func (spq *SimplePriorityQueue) Update(task *SimpleTask, dbID int64, value string, priority int) {
	task.DBID = dbID
	task.Value = value
	task.Priority = priority
	heap.Fix(spq, task.index)
}
