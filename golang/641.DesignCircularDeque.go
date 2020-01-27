type MyCircularDeque struct {
    queue []int
	capicity int
	front, rear int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	var queue MyCircularDeque
	queue.queue = make([]int, k + 1)
	queue.capicity = k + 1
	return queue
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.front = (this.front - 1 + this.capicity) % this.capicity
	this.queue[this.front] = value
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
		return false
	}

	this.queue[this.rear] = value
	this.rear = (this.rear + 1) % this.capicity
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
		return false
	}

	this.front = (this.front + 1) % this.capicity
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
		return false
	}

	this.rear = (this.rear - 1 + this.capicity) % this.capicity
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
		return -1
	}

	return this.queue[this.front]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
		return -1
	}

	return this.queue[(this.rear - 1 + this.capicity) % this.capicity]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    return this.front == this.rear
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    return (this.rear + 1) % this.capicity == this.front
}

