package iterator

type Iterator interface {
	HasNext() bool
	Next() int
}

type AbstractAggregate interface {
	createIterator() Iterator
}

type ArrayList struct {
	Data []int
}

func (a *ArrayList) createIterator() Iterator {
	return &ArrayListIterator{List: a, Index: 0}
}

// Iterate through an array sequentially
type ArrayListIterator struct {
	List  *ArrayList
	Index int
}

func (i *ArrayListIterator) HasNext() bool {
	return i.Index < len(i.List.Data)
}

func (i *ArrayListIterator) Next() int {
	if i.HasNext() {
		i.Index += 1
		return i.List.Data[i.Index-1]
	}
	return -1
}

// Iterate through an array while skipping duplicate elements

type UniqueArrayIterator struct {
	List    *ArrayList
	Index   int
	visited map[int]bool
}

func (i *UniqueArrayIterator) HasNext() bool {
	l := len(i.List.Data)
	for idx := i.Index; idx < l; idx++ {
		if _, ok := i.visited[i.List.Data[idx]]; !ok {
			return true
		}
	}
	return false
}

func (i *UniqueArrayIterator) Next() int {
	if i.HasNext() {
		value := i.List.Data[i.Index]
		i.Index++
		l := len(i.List.Data)
		for idx := i.Index; idx < l; idx++ {
			if _, ok := i.visited[i.List.Data[idx]]; !ok {
				i.visited[i.List.Data[idx]] = true
				i.Index = idx
			}
		}
		return value
	}
	return -1
}

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func (l *LinkedList) createIterator(a AbstractAggregate) Iterator {
	return &LinkedListIterator{CurrentNode: l}
}

type LinkedListIterator struct {
	CurrentNode *LinkedList
}

func (l *LinkedListIterator) HasNext() bool {
	return l.CurrentNode != nil
}

func (l *LinkedListIterator) Next() int {
	if l.HasNext() {
		value := l.CurrentNode.Value
		l.CurrentNode = l.CurrentNode.Next
		return value
	}
	return -1
}

type TreeNode struct {
	value      int
	childNodes []*TreeNode
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) createIterator(a AbstractAggregate) Iterator {
	return nil
}

type TreeIterator interface {
	HasNext() bool
	Next() int
}

func CreateIterator(a AbstractAggregate) Iterator {
	return a.createIterator()
}
