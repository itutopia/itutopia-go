package listes

import (
	"container/list"
	"fmt"
)

// list
/**
Go语言内置容器list是一个双向链表(实际是一个环)。位于包list当中。
list的核心结构体一共包含两个List和Element。
结构体：
	type List struct {
		root Element // sentinel list element, only &root, root.prev, and root.next are used
		len  int     // current list length excluding (this) sentinel element
	}

List中Element的结构体如下：
type Element struct {
    list *List
    Value interface{}
       next, prev *Element
}
list：指向List的指针，用于标识当前节点属于哪一个LIST
Value：用于存储元素的值。
next, prev： 指向Element的指针，用于定位后一个节点与前一个节点。
*/
// 官方
func ListExample() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// 遍历输出链表内容
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// list模拟栈结构
func ListStackExample() {
	stack := list.New()
	stack.PushBack(1)
	stack.PushBack(2)
	stack.PushBack(3)
	stack.PushBack(4)
	fmt.Println("栈顶元素的值:", stack.Back().Value)

	for stack.Len() > 0 {
		fmt.Println(stack.Remove(stack.Back()))
	}
}

// list模拟队列结构
func ListQueueExample() {
	queue := list.New()
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)
	queue.PushBack(4)

	for queue.Len() > 0 {
		fmt.Println(queue.Remove(queue.Front()))
	}
}

func ListApplicationExample() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	fmt.Println(l.Len())

	// 遍历输出链表内容
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
