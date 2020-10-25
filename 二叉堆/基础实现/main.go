package main

type MaxPQ struct {
	pq []int
	n  int
}

func (m *MaxPQ) max() int {
	if len(m.pq) <= 1 {
		return -1
	}
	return m.pq[1]
}

func (m *MaxPQ) insert(n int) {
}

func (m *MaxPQ) delMax() int {
	if len(m.pq) <= 1 {
		return -1
	}
	return m.pq[1]
}

func (m *MaxPQ) swim(k int) {
	for k > 1 && m.less(m.parent(k), k) {
		m.exchange(m.parent(k), k)
		k = m.parent(k)
	}
}

func (m *MaxPQ) less(i int, j int) bool {
	return m.pq[i] < m.pq[j]
}

func (m *MaxPQ) sink(k int) {
	for m.left(k) <= len(m.pq) {
		olderIndex := m.left(k)
		if m.right(k) <= len(m.pq) && m.less(m.left(k), m.right(k)) {
			olderIndex = m.right(k)
		}

		if m.less(olderIndex, k) {
			break
		}

		m.exchange(olderIndex, k)
		k = olderIndex
	}
}

func (m *MaxPQ) exchange(i int, j int) {
	temp := m.pq[i]
	m.pq[i] = m.pq[j]
	m.pq[j] = temp
}

func newMaxPQ() *MaxPQ {
	return &MaxPQ{
		pq: []int{-1},
		n:  0,
	}
}

func (m *MaxPQ) parent(root int) int {
	return root / 2
}

func (m *MaxPQ) left(root int) int {
	return root * 2
}

func (m *MaxPQ) right(root int) int {
	return root*2 + 1
}

func main() {

}
