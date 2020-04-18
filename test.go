package main

import (
	"fmt"
	"sync"
	"time"
)

var list List

func main() {
	var arr [30000000]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	for i := 0; i < 30000000; i++ {
		list.RPush(i + 1)
	}
	start := time.Now()

	a := make(chan int, 1000)
	group := len(arr) / 500000
	mod := len(arr) % 500000
	var count int
	if mod == 0 {
		count = group
	} else {
		count = group + 1
	}
	wg := sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		end := 500000 * (i + 1)
		if end < len(arr) {
			go add(arr[end-500000:end], &wg, a)
		} else {
			go add(arr[end-500000:len(arr)], &wg, a)
		}
	}
	wg.Wait()
	close(a)
	var sum int
	for i := range a {
		sum += i
	}
	fmt.Println(sum)

	cost := time.Since(start)
	fmt.Println(cost)

	btime := time.Now()
	addsec(arr[:])
	bcost := time.Since(btime)
	fmt.Println(bcost)

	ctime := time.Now()
	cb := make(chan int, 1000)
	wg2 := sync.WaitGroup{}

	for i := 0; i < 60; i++ {
		wg2.Add(1)
		go add2(&wg2, cb)
	}
	wg2.Wait()
	close(cb)
	var sum2 int
	for i := range a {
		sum2 += i
	}
	fmt.Println(sum2)

	ccost := time.Since(ctime)
	fmt.Printf("ccost: %d\n", ccost)

}

func add(arr []int, wg *sync.WaitGroup, a chan int) {
	var sum int
	for _, i := range arr {
		sum += i
	}
	wg.Done()
	a <- sum
}

func add2(wg *sync.WaitGroup, a chan int) {
	var sum int
	for {
		b := list.LPop()
		if b == nil {
			break
		}
		sum += b.GetValue()
		//fmt.Println(sum)
	}
	wg.Done()
	a <- sum
}

/*func findData() int {
	if len(arr) > 0 {
		result := arr[0]
		arr = arr[1:]
		//fmt.Println(result)
		return result
	} else {
		return 0
	}
}*/

func addsec(arr []int) {
	var sum int
	for _, i := range arr {
		sum += i
	}
	fmt.Println(sum)
}

////////////////////////////////////////////////////////////////////////
// 链表的一个节点
type ListNode struct {
	prev  *ListNode // 前一个节点
	next  *ListNode // 后一个节点
	value int       // 数据
}

// 创建一个节点
func NewListNode(value int) (listNode *ListNode) {
	listNode = &ListNode{
		value: value,
	}

	return
}

// 当前节点的前一个节点
func (n *ListNode) Prev() (prev *ListNode) {
	prev = n.prev

	return
}

// 当前节点的前一个节点
func (n *ListNode) Next() (next *ListNode) {
	next = n.next

	return
}

// 获取节点的值
func (n *ListNode) GetValue() (value int) {
	if n == nil {

		return
	}
	value = n.value

	return
}

////////////////////////////////////////////////////////////////////////
// 链表
type List struct {
	head *ListNode // 表头节点
	tail *ListNode // 表尾节点
	len  int       // 链表的长度
}

// 创建一个空链表
func NewList() (list *List) {
	list = &List{}
	return
}

// 返回链表头节点
func (l *List) Head() (head *ListNode) {
	head = l.head

	return
}

// 返回链表尾节点
func (l *List) Tail() (tail *ListNode) {
	tail = l.tail

	return
}

// 返回链表长度
func (l *List) Len() (len int) {
	len = l.len

	return
}

////////////////////////////////////////////////////////////////////////
// 在链表的右边插入一个元素
func (l *List) RPush(value int) {

	node := NewListNode(value)

	// 链表未空的时候
	if l.Len() == 0 {
		l.head = node
		l.tail = node
	} else {
		tail := l.tail
		tail.next = node
		node.prev = tail

		l.tail = node
	}

	l.len = l.len + 1

	return
}

////////////////////////////////////////////////////////////////////////
// 从链表左边取出一个节点
func (l *List) LPop() (node *ListNode) {

	// 数据为空
	if l.len == 0 {

		return
	}

	node = l.head

	if node.next == nil {
		// 链表未空
		l.head = nil
		l.tail = nil
	} else {
		l.head = node.next
	}
	l.len = l.len - 1

	return
}

////////////////////////////////////////////////////////////////////////
