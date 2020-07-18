package unionfind

<<<<<<< HEAD
// Robert Sedgewick 算法（第4版） 1.5.2.7
// union-find (加权 quick-union)，还作了路径压缩优化

// union is ...
type union struct {
	id    []int // 父链接数组(由触点索引)
	sz    []int // (由触点索引的) 各个根节点所对应的分量的大小
	count int   // 连通分量的数量
}

func newUnion(N int) *union {
	id := make([]int, N)
	for i := range id {
		id[i] = i
	}
	sz := make([]int, N)
	for i := range sz {
		sz[i] = 1
	}
	return &union{
		id:    id,
		sz:    sz,
		count: N,
	}
}

func (u *union) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *union) find(p int) int {
	// 跟随连接找到根节点
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}

func (u *union) union(p, q int) {
	i, j := u.find(p), u.find(q)
	if i == j {
		return
	}
	if u.sz[i] > u.sz[j] {
		i, j = j, i
	}
	// 将小树的根节点连接到大树的根节点
	u.id[i] = j
	u.sz[j] += u.sz[i]
	u.count--
	return
=======
// leetcode 中需要 union 的集合都太小了，最普通的集合算法，就够用了

type unionFind struct {
	parent []int
}

func newUnionFind(size int) *unionFind {
	parent := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return &unionFind{
		parent: parent,
	}
}

func (u *unionFind) find(i int) int {
	if u.parent[i] != i {
		u.parent[i] = u.find(u.parent[i])
	}
	return u.parent[i]
}

func (u *unionFind) union(x, y int) {
	u.parent[u.find(x)] = u.find(y)
>>>>>>> f33c3a477711033e1c5c5c04e72ce2c3c83f449e
}
