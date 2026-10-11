type MyHashMap struct {
	data []int
}

func Constructor() MyHashMap {
    return MyHashMap{
		data: make([]int, 0),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	if len(this.data) < key+1 {
		for i := len(this.data); i < key+1; i++ {
			this.data = append(this.data, -1)
		}
	}
	this.data[key] = value
}

func (this *MyHashMap) Get(key int) int {
	if len(this.data) < key+1 {
		return -1
	}
	return this.data[key]
}

func (this *MyHashMap) Remove(key int) {
	if len(this.data) > key {
		this.data[key] = -1
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */