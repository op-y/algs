package uf

type UF struct {
	count int
	id    []int
}

func NewUF(n int) {
	u := new(UF)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	u.count = n
	u.id = id
	return u
}

func (u *UF) find(p int) int {
	return u.id[p]
}

func (u *UF) Count() int {
	return u.count
}

func (u *UF) Connected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *UF) Union(p, q int) {
	pID := u.find(p)
	qID := u.find(q)

	if pID == qID {
		return
	}

	for idx, value := range u.id {
		if value == pID {
			u.id[idx] = qID
		}
	}

	u.count--
	return
}
