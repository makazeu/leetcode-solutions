type MyCircularQueue struct {
	queue []int
	capicity int
	front, rear int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	var queue MyCircularQueue
	queue.queue = make([]int, k + 1)
	queue.capicity = k + 1
	return queue
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
    if (this.rear + 1) % this.capicity == this.front {
		return false
	}
	this.queue[this.rear] = value
	this.rear = (this.rear + 1) % this.capicity
    return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
    if this.front == this.rear {
		return false
	}
	this.front = (this.front + 1) % this.capicity
    return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
    if this.front == this.rear {
		return -1
	}
	return this.queue[this.front]
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
    if this.front == this.rear {
		return -1
	}
    return this.queue[(this.rear + this.capicity - 1) % this.capicity]
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
    return this.front == this.rear
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
    return (this.rear + 1) % this.capicity == this.front
}

