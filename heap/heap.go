// Heap implementation from The Algorithm Design Manual by Steven S. Skiena.
package heap

type Heap struct {
	queue []int
}

func (h *Heap) Insert(n int) {
	h.queue = append(h.queue, n)
	bubbleUp(h.queue, len(h.queue)-1)
}

func (h *Heap) ExtractMin() int {
	min := -1

	if len(h.queue) > 0 {
		min = h.queue[0]

		h.queue[0] = h.queue[len(h.queue)-1]
		h.queue = h.queue[:len(h.queue)-1]
		bubbleDown(h.queue, 0)
	}

	return min
}

func bubbleUp(queue []int, i int) {
	if parent(i) == -1 {
		return
	} else if queue[parent(i)] > queue[i] {
		swap(queue, parent(i), i)
		bubbleUp(queue, parent(i))
	}
}

func bubbleDown(queue []int, i int) {
	childIndex := firstChild(i)
	lightestIndex := i

	for i := 0; i <= 1; i++ {
		if (childIndex + i) <= (len(queue) - 1) {
			if queue[lightestIndex] > queue[childIndex+i] {
				lightestIndex = childIndex + i
			}
		}
	}

	if lightestIndex != i {
		swap(queue, i, lightestIndex)
		bubbleDown(queue, lightestIndex)
	}
}

func parent(i int) int {
	if i == 0 {
		return -1
	} else {
		return (i - 1) / 2
	}
}

func firstChild(i int) int {
	return (2 * i) + 1
}

func swap(queue []int, i int, j int) {
	queue[i], queue[j] = queue[j], queue[i]
}
