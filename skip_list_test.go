package easy_go

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_SkipList(t *testing.T) {

	n := 100
	rand.Seed(time.Now().Unix())

	// 向跳表中加入100个0~10000之间的数
	skipList := NewSkipList()
	for skipList.Size() < n {
		skipList.Put(NewInteger(rand.Intn(10000)))
	}
	skipList.PrintInfo()

	// 查询跳表中值在100以内的数
	fmt.Print("these numbers will be deleted: ")
	for i := 0; i < n; i++ {
		get := skipList.Get(NewInteger(i))
		if get != nil {
			fmt.Printf("%d ", get.(Integer).value)
		}
	}
	fmt.Println()

	// 删除跳表中值在100以内的数
	for i := 0; i < n; i++ {
		skipList.Remove(NewInteger(i))
	}
	skipList.PrintInfo()

}
