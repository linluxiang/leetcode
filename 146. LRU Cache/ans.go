import "container/list"

type LRUCache struct {
    cap int
    data map[int]*list.Element // from key to the element
    lst *list.List // the list is arranged as the first one is least recently used
}

type Entry struct {
    Key int
    Val int
}

func Constructor(capacity int) LRUCache {
    return LRUCache {
        cap: capacity,
        data: map[int]*list.Element{},
        lst: list.New(),
    }
}


func (this *LRUCache) Get(key int) int {
    if ele, ok := this.data[key]; ok {
        this.lst.MoveToBack(ele)
        v := ele.Value.(*Entry)
        return v.Val
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if ele, ok := this.data[key]; ok {
        this.lst.MoveToBack(ele)
        v := ele.Value.(*Entry)
        v.Val = value
        return
    }
    // check cap
    if this.cap == len(this.data) {
        front := this.lst.Front()
        v := front.Value.(*Entry)
        delete(this.data, v.Key)
        this.lst.Remove(front)
    }

    this.lst.PushBack(&Entry{key, value})
    ele := this.lst.Back()
    this.data[key] = ele
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
