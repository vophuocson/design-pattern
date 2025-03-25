package iterator

import (
	"design-pattern-practice/iterator/iterator"
	"fmt"
)

func main() {
	arr := iterator.ArrayList{
		Data: []int{1, 2, 3, 4},
	}
	it := iterator.CreateIterator(&arr)
	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
