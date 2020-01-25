type Node struct {
	key   int
	value int
	count int
}

type LFUCache struct {
	capacity  int
	list      *list.List
	hash      map[int]*list.Element
	countHash map[int]*list.Element
}

func Constructor(capacity int) LFUCache {
	var cache LFUCache
	cache.capacity = capacity
	cache.list = list.New()
	cache.hash = make(map[int]*list.Element)
	cache.countHash = make(map[int]*list.Element)
	return cache
}

func (this *LFUCache) updateCount(element *list.Element) {
	count := element.Value.(*Node).count

	if curElement := this.countHash[count]; curElement == element {
		if curElement != this.list.Front() && curElement.Prev().Value.(*Node).count == count {
			this.countHash[count] = curElement.Prev()
		} else {
			delete(this.countHash, count)
		}
	} else {
		this.list.MoveAfter(element, curElement)
	}

	count++
	element.Value.(*Node).count = count
	if nxtElement, exists := this.countHash[count]; exists {
		this.list.MoveAfter(element, nxtElement)
	}
	this.countHash[count] = element
}

func (this *LFUCache) Get(key int) int {
	if element, exists := this.hash[key]; exists {
		this.updateCount(element)
		return element.Value.(*Node).value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if element, exists := this.hash[key]; exists {
		element.Value.(*Node).value = value
		this.updateCount(element)
		return
	}

	if this.list.Len() >= this.capacity && this.list.Len() > 0 {
		count := this.list.Front().Value.(*Node).count
		if firstElement := this.countHash[count]; firstElement == this.list.Front() {
			delete(this.countHash, count)
		}

		delete(this.hash, this.list.Front().Value.(*Node).key)
		this.list.Remove(this.list.Front())
	}

	if this.list.Len() >= this.capacity {
		return
	}

	node := &Node{key: key, value: value, count: 1}
	var newElement *list.Element
	if curElement, exists := this.countHash[node.count]; exists {
		newElement = this.list.InsertAfter(node, curElement)
	} else {
		newElement = this.list.PushFront(node)
	}
	this.hash[key] = newElement
	this.countHash[node.count] = newElement
}

