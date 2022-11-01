package easy

type MyHashMap struct {
	biggestIDX int
	keyvals    []int
}

func Constructor1() MyHashMap {
	myHm := MyHashMap{
		biggestIDX: 1000000,
		keyvals:    make([]int, 1000000),
	}

	for i := 0; i < len(myHm.keyvals); i++ {
		myHm.keyvals[i] = -1
	}

	return myHm
}

func (h *MyHashMap) Put(key int, value int) {
	if key > h.biggestIDX {
		h.biggestIDX = key + 1
		newKeyVals := make([]int, h.biggestIDX)
		copy(newKeyVals, h.keyvals)
		h.keyvals = newKeyVals
	}

	h.keyvals[key] = value
}

func (h *MyHashMap) Get(key int) int {
	return h.keyvals[key]
}

func (h *MyHashMap) Remove(key int) {
	h.keyvals[key] = -1
}
