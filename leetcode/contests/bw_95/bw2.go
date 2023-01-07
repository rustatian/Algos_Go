package bw_95

type DataStream struct {
	value    int
	lastInts []int
	k        int
}

func Constructor(value int, k int) DataStream {
	return DataStream{
		value:    value,
		lastInts: make([]int, 0, k),
		k:        k,
	}
}

func (ds *DataStream) Consec(num int) bool {
	ds.lastInts = append(ds.lastInts, num)

	if len(ds.lastInts) < ds.k {
		return false
	}

	if len(ds.lastInts) == ds.k {
		for i := 0; i < len(ds.lastInts); i++ {
			if ds.lastInts[i] != ds.value {
				ds.lastInts = ds.lastInts[1:]
				return false
			}
		}

		// remove first element
		ds.lastInts = ds.lastInts[1:]
		return true
	}

	return false
}
