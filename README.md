# easy-go
Data structure and algorithm implementation for golang

## usage
### 1.PriorityQueue
Define your struct and implements 'easy-go/Interface'. 

Actually you only need to implement 'CompareTo(e interface{}) int {}'.
```go
type Person struct {
    name string
    age  int
}

// CompareTo compare by person.age then person.name
func (p Person) CompareTo(e interface{}) int {
    if p.age == e.(Person).age {
        return strings.Compare(p.name, e.(Person).name)
    }
    return p.age - e.(Person).age
}

func exampleOfPriorityQueue() {
    // use 'easy_go.Interface' to make a slice instead of 'Person'
    slice := make([]easy_go.Interface, 0)
    slice = append(slice, Person{"Alice", 18})
    slice = append(slice, Person{"Bob", 17})
    slice = append(slice, Person{"John", 18})
    slice = append(slice, Person{"Mark", 19})
    // you can get a PriorityQueue by inputting a slice
    priorityQueue := easy_go.NewPriorityQueueWithSlice(slice)
    // push
    priorityQueue.Push(Person{"Bill", 16})
    for priorityQueue.Size() != 0 {
        // pop
        fmt.Println(priorityQueue.Pop())
    }
}
```

### 2.SkipList
Define your struct and implements 'easy-go/Interface'.

Actually you only need to implement 'CompareTo(e interface{}) int {}'.

You can implement a variety of comparison functions by passing function pointers to the constructor.
```go
func exampleOfSkipList() {
    // define a unnamed compare-function to replace 'person.CompareTo'
    skipList := easy_go.NewSkipListWithCompare(func(a, b easy_go.Interface) int {
        return strings.Compare(a.(Person).name, b.(Person).name)
    })
    
    skipList.Put(Person{"Alice", 18})
    skipList.Update(Person{"Alice", 19})
    skipList.Update(Person{"Bob", 17})
    skipList.Put(Person{"Bob", 15})
    skipList.Put(Person{"Mark", 19})
    skipList.Remove(Person{"Mark", -1})

    Alice := skipList.Get(Person{"Alice", -1})
    Bob := skipList.Get(Person{"Bob", -1})
    Mark := skipList.Get(Person{"Mark", -1})
    
    fmt.Println(Alice)
    fmt.Println(Bob)
    fmt.Println(Mark)
}
```
### 2.BitMap
```go
func exampleOfBitMap() {
	bitMap := easy_go.NewBitMap(33)
	bitMap.SetTrue(0)
	bitMap.SetTrue(1)
	bitMap.SetTrue(32)
	bitMap.SetTrue(31)
	bitMap.PrintInfo()
	bitMap.SetFalse(31)
	bitMap.SetFalse(0)
	fmt.Println(bitMap.Get(1))
	fmt.Println(bitMap.Get(0))
	bitMap.PrintInfo()
}
```

