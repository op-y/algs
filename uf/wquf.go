package uf

type WeightedQuickUF struct {
	count  int
	id     []int
	weight []int
}

func NewWeightedQuickUF(n int) *WeightedQuickUF {
	u := new(WeightedQuickUF)
	id := make([]int, n)
	weight := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
		weight[i] = 1
	}
	u.count = n
	u.id = id
	u.weight = weight
	return u
}

func (u *WeightedQuickUF) find(p int) int {
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}

func (u *WeightedQuickUF) Count() int {
	return u.count
}

func (u *WeightedQuickUF) Connected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *WeightedQuickUF) Union(p, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)
	if pRoot == qRoot {
		return
	}
	if u.weight[pRoot] < u.weight[qRoot] {
		u.id[pRoot] = qRoot
		u.weight[qRoot] += u.weight[pRoot]
	} else {
		u.id[qRoot] = pRoot
		u.weight[pRoot] += u.weight[qRoot]
	}
	u.count--
	return
}
