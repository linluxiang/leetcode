package main

import (
	"container/list"
	"fmt"
)

type Entry struct {
	Val int
	Idx int
}

type MonotonicQueue struct {
	queue  *list.List
	length int
}

func NewMonotonicQueue(length int) *MonotonicQueue {
	return &MonotonicQueue{
		queue:  list.New(),
		length: length,
	}
}

func (q *MonotonicQueue) Push(entry *Entry) {
	if q.queue.Back() == nil {
		q.queue.PushBack(entry)
		return
	}
	back := q.queue.Back()
	backValue := back.Value.(*Entry)
	for backValue.Val > entry.Val {
		q.queue.Remove(back)
		back = q.queue.Back()
		if back == nil {
			break
		}
		backValue = back.Value.(*Entry)
	}
	q.queue.PushBack(entry)
	if q.queue.Len() > q.length {
		q.queue.Remove(q.queue.Front())
	}

	front := q.queue.Front()
	frontValue := front.Value.(*Entry)
	for frontValue.Idx <= entry.Idx-q.length {
		q.queue.Remove(front)
		front := q.queue.Front()
		if front == nil {
			break
		}
		frontValue = front.Value.(*Entry)
	}
}

func (q *MonotonicQueue) Front() *Entry {
	if q.queue.Front() == nil {
		return nil
	}

	return q.queue.Front().Value.(*Entry)
}

func (q *MonotonicQueue) Len() int {
	return q.queue.Len()
}

func maxSubarraySumCircular(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	wsize := len(nums)

	for i := 0; i < wsize-1; i++ {
		nums = append(nums, nums[i])
	}

	queue := NewMonotonicQueue(wsize)
	queue.Push(&Entry{nums[0], 0})

	sum := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		try := sum - queue.Front().Val
		if try > ans {
			ans = try
		}
		queue.Push(&Entry{sum, i})
	}

	return ans
}

func main() {
	ans := maxSubarraySumCircular([]int{3, -1, 2, -1})
	fmt.Println(ans)
}
