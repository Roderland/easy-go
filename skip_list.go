package easy_go

import (
	"fmt"
	"math/rand"
	"time"
)

const MaxLevel = 32
const LevelFactor = 0.25

type (
	SkipList struct {
		head  node
		level int
		size  int
		cmp   func(a, b Interface) int
	}

	node struct {
		data Interface
		next []*node
	}
)

func newNode(data Interface, lev int) *node {
	return &node{
		data: data,
		next: make([]*node, lev),
	}
}

func NewSkipList() *SkipList {
	return &SkipList{
		head:  *newNode(nil, MaxLevel),
		level: 0,
		size:  0,
	}
}

func NewSkipListWithCompare(cmp func(a, b Interface) int) *SkipList {
	return &SkipList{
		head:  *newNode(nil, MaxLevel),
		level: 0,
		size:  0,
		cmp:   cmp,
	}
}

func (sp *SkipList) compare(a, b Interface) int {
	if sp.cmp != nil {
		return sp.cmp(a, b)
	} else {
		return a.CompareTo(b)
	}
}

func (sp *SkipList) Get(data Interface) Interface {
	pre := &sp.head
	for i := sp.level - 1; i >= 0; i-- {
		pre = sp.findClosest(pre, i, data)
		cur := pre.next[i]
		if cur != nil && sp.compare(data, cur.data) == 0 {
			return cur.data
		}
	}
	return nil
}

func (sp *SkipList) Size() int {
	return sp.size
}

func (sp *SkipList) IsExist(data Interface) bool {
	return sp.Get(data) != nil
}

// Put insert only not exist
func (sp *SkipList) Put(data Interface) bool {
	if sp.IsExist(data) {
		return false
	}
	lev := sp.randomLevel()
	nnd := newNode(data, lev)
	pre := &sp.head
	if lev <= sp.level {
		for i := lev - 1; i >= 0; i-- {
			pre = sp.findClosest(pre, i, nnd.data)
			cur := pre.next[i]
			pre.next[i] = nnd
			nnd.next[i] = cur
		}
	} else {
		for i := lev - 1; i >= sp.level; i-- {
			pre.next[i] = nnd
		}
		for i := sp.level - 1; i >= 0; i-- {
			pre = sp.findClosest(pre, i, nnd.data)
			cur := pre.next[i]
			pre.next[i] = nnd
			nnd.next[i] = cur
		}
		sp.level = lev
	}
	sp.size++
	return true
}

// Update insert only exist
func (sp *SkipList) Update(data Interface) bool {
	if !sp.IsExist(data) {
		return false
	}
	pre := &sp.head
	for i := sp.level - 1; i >= 0; i-- {
		pre = sp.findClosest(pre, i, data)
		cur := pre.next[i]
		if cur != nil && sp.compare(data, cur.data) == 0 {
			cur.data = data
		}
	}
	return true
}

func (sp *SkipList) Remove(data Interface) bool {
	res := false
	pre := &sp.head
	for i := sp.level - 1; i >= 0; i-- {
		pre = sp.findClosest(pre, i, data)
		cur := pre.next[i]
		if cur != nil && sp.compare(data, cur.data) == 0 {
			pre.next[i] = cur.next[i]
			res = true
		}
	}
	if res {
		sp.size--
	}
	return res
}

func (sp *SkipList) PrintInfo() {
	fmt.Printf("level: %d, size: %d\n", sp.level, sp.size)
	cur := sp.head.next[0]
	idx := 0
	for cur != nil {
		fmt.Printf("[%d: %v] ", idx, cur)
		idx++
		cur = cur.next[0]
	}
	fmt.Println()
}

func (sp *SkipList) findClosest(cur *node, lev int, data Interface) *node {
	for cur.next[lev] != nil && sp.compare(data, cur.next[lev].data) > 0 {
		cur = cur.next[lev]
	}
	return cur
}

func (sp *SkipList) randomLevel() int {
	lev := 1
	rand.Seed(time.Now().Unix())
	for rand.Float32() < LevelFactor && lev <= sp.level && lev < MaxLevel {
		lev++
	}
	return lev
}
