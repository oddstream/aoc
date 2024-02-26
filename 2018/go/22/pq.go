package main

// https://pkg.go.dev/container/heap

type PQItem struct {
	Point    // can use this to get Region.risk/type from cave map
	equipped EquipmentType
	minutes  int
	index    int // internal use by heap
}

type PriorityQueue []*PQItem

// need to implement sort interface (Len, Less, Swap)
// and heap interface (Push, Pop)

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].minutes == pq[j].minutes {
		return pq[i].manhatten() < pq[j].manhatten()
	}
	return pq[i].minutes < pq[j].minutes
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*PQItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
// func (pq *PriorityQueue) update(item *Item, point Point, priority int) {
// 	item.Point = point
// 	item.risk = priority
// 	heap.Fix(pq, item.index)
// }
