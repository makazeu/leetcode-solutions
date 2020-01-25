type Node struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	list     *list.List
	hash     map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	var cache LRUCache
	cache.capacity = capacity
	cache.list = list.New()
	cache.hash = make(map[int]*list.Element)
	return cache
}

func (this *LRUCache) Get(key int) int {
	if element, exists := this.hash[key]; exists {
		this.list.MoveToBack(element)
		return element.Value.(*Node).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if element, exists := this.hash[key]; exists {
		this.list.MoveToBack(element)
		element.Value.(*Node).value = value
		return
	}

	if this.list.Len() >= this.capacity {
		delete(this.hash, this.list.Front().Value.(*Node).key)
		this.list.Remove(this.list.Front())
	}

	node := &Node{key: key, value: value}
	this.hash[key] = this.list.PushBack(node)
}

