package easy_go

import "fmt"

type BitMap []byte

func NewBitMap(size int) BitMap {
	length := size / 8
	if size%8 != 0 {
		length++
	}
	return make([]byte, length)
}

func (bm *BitMap) Get(index int) bool {
	i := index / 8
	j := index % 8
	return ((*bm)[i] & (1 << j)) != 0
}

func (bm *BitMap) SetTrue(index int) {
	i := index / 8
	j := index % 8
	(*bm)[i] |= 1 << j
}

func (bm *BitMap) SetFalse(index int) {
	i := index / 8
	j := index % 8
	(*bm)[i] &= ^(1 << j)
}

func (bm *BitMap) PrintInfo() {
	fmt.Printf("length of bytes: %d\n", len(*bm))
	for i, b := range *bm {
		fmt.Printf("index: %d, value: %08b\n", i, b)
	}
}
