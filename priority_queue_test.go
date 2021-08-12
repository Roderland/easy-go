package easy_go

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_PriorityQueue(t *testing.T) {

	n := 100
	rand.Seed(time.Now().Unix())

	slice := make([]Interface, n)
	for i := 0; i < n; i++ {
		slice[i] = NewInteger(rand.Intn(100))
	}
	pq := NewPriorityQueueWithSlice(slice)

	pq.PrintInfo()
	for i := 0; pq.size != 0; i++ {
		slice[i] = pq.Pop()
	}
	fmt.Println(slice)
}
