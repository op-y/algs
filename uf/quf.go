package uf

type QuickUF struct {
	count int
	id    []int
}

func NewQuickUF(n int) *QuickUF {
	u := new(QuickUF)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	u.count = n
	u.id = id
	return u
}

func (u *QuickUF) find(p int) int {
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}

func (u *QuickUF) Count() int {
	return u.count
}

func (u *QuickUF) Connected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *QuickUF) Union(p, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)
	if pRoot == qRoot {
		return
	}
	u.id[pRoot] = qRoot
	u.count--
	return
}
