type Item struct {
	data int
	min  int
}

type MinStack struct {
	*list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{list.New()}
}

func (this *MinStack) Push(x int) {
	item := &Item{x, x}

	if this.Len() > 0 {
		min := this.Back().Value.(*Item).min
		if min < x {
			item.min = min
		}
	}
    
    this.PushBack(item)
}

func (this *MinStack) Pop() {
	this.Remove(this.Back())
}

func (this *MinStack) Top() int {
	return this.Back().Value.(*Item).data
}

func (this *MinStack) GetMin() int {
	return this.Back().Value.(*Item).min
}

