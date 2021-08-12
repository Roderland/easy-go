# easy-go
Data structure and algorithm implementation for golang

## usage
### 1.PriorityQueue
Define your struct and implements 'easy-go/Interface'. 

Actually you only need to implement 'func Less(e interface{}) bool {}'.
```go
type Person struct {
    name string
    age int
}

// Less compare by person.age
func (p Person) Less(e interface{}) bool {
    return p.age < e.(Person).age
}

func main() {
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
