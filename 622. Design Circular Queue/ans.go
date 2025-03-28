type Node struct {
    Val int
    Next *Node
}

type MyCircularQueue struct {
    head *Node
    tail *Node
    cap int
    length int
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
        head: nil,
        tail: nil,
        cap: k,
        length: 0,
    }
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.length == this.cap {
        return false
    }

    node := &Node{
        Val: value,
        Next: nil,
    }

    if this.tail == nil {
        // empty
        this.head = node
        this.tail = node
    } else {
        this.tail.Next = node
        this.tail = this.tail.Next
    }
    

    this.length += 1
    this.tail.Next = this.head
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    // if there is only one 
    // set the head and tail to nil
    // else head = head.next
    // tail.next to head

    if this.length == 0 {
        return false
    }

    if this.length == 1 {
        this.head = nil
        this.tail = nil
        this.length = 0
        return true
    }

    this.head = this.head.Next
    this.tail.Next = this.head
    this.length -= 1
    return true
}


func (this *MyCircularQueue) Front() int {
    if this.head == nil {
        return -1
    }
    return this.head.Val
}


func (this *MyCircularQueue) Rear() int {
    if this.tail == nil {
        return -1
    }

    return this.tail.Val
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.length == 0
}


func (this *MyCircularQueue) IsFull() bool {
    return this.length == this.cap
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
