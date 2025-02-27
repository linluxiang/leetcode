import "container/list"

func rotate(nums []int, k int)  {
    l := list.New()
    for _, num := range nums {
        l.PushBack(num)
    }

    for range k {
        l.MoveToFront(l.Back())
    }

    i := 0
    head := l.Front()

    for head != nil {
        nums[i] = head.Value.(int)
        head = head.Next()
        i++
    }

}
