package main 

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {

		r = append(r, k)
	}
	return r
}


type List[T any] struct {
	head tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

func (l *List[T]) Push(val T) {
	if lst.tail == nil {
		lst.tail = &element[T]{val: V}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: V}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) ForEach(f func(T)) {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("keys:", MapKeys(m))
    

	_= Mapkeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())	
}