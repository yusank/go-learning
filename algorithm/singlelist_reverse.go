package main

import (
	"fmt"
)

type List struct {
	first *element
	last  *element
	size  int
}

type element struct {
	value interface{}
	next  *element
}

func newList(sli []string) *List {
	l := &List{}

	for _, v := range sli {
		newelement := &element{value: v}
		if l.size == 0 {
			l.first = newelement
			l.last = newelement
		} else {
			l.last.next = newelement
			l.last = newelement
		}
		l.size++
	}

	return l
}

func (list *List) reverse() {
	var p *element

	if list.size == 0 {
		fmt.Print("nil list")
		return
	}

	p = list.first

	for ele := list.first.next; ele != nil; ele = ele.next {
		q := *ele
		p.next = q.next
		q.next = list.first
		list.first = &q
	}

	after := make([]interface{}, list.size, list.size)
	for e, elem := 0, list.first; elem != nil; e, elem = e+1, elem.next {
		after[e] = elem.value
	}

	fmt.Printf("after reverse: %v \n", after)

}

func main() {
	s := []string{"a", "b", "c", "d", "f"}

	l := newList(s)

	fmt.Printf("before reverse: %v \n", s)
	l.reverse()
}
