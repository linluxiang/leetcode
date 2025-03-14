
import (
	"fmt"
    "container/heap"
)

type Entry struct {
    Id int
    Score int
    Idx int
}

type Queue []*Entry

func (q *Queue) Push(v any) {
    entry := v.(*Entry)

    *q = append(*q, entry)
    entry.Idx = len(*q) - 1
}

func (q *Queue) Pop() any {
    old := *q
    size := len(old)
    if size == 0 {
        return nil
    }
    last := old[size - 1]
    old[size - 1] = nil
    last.Idx = -1
    *q = old[:size - 1]
    return last
}

func (q Queue) Len() int {
    return len(q)
}

func (q Queue) Less(a, b int) bool {
    return q[b].Score < q[a].Score
}

func (q Queue) Swap(i, j int) {
    q[i], q[j] = q[j], q[i]
    q[i].Idx = i
    q[j].Idx = j
}

func (q *Queue) Update(entry *Entry, score int) {
    entry.Score = score
    heap.Fix(q, entry.Idx)
}

type Leaderboard struct {
    entries *Queue
    playerIdx map[int]*Entry // id to idx
    k int
}


func Constructor() Leaderboard {
    return Leaderboard{
        entries: &Queue{},
        playerIdx: map[int]*Entry{},
    }
}


func (this *Leaderboard) AddScore(playerId int, score int)  {
    // entry := &Entry{
    //     Id: playerId,
    //     Score: score,
    // }
    if entry, ok := this.playerIdx[playerId]; ok {
        this.entries.Update(entry, entry.Score + score)
    } else {
        entry := &Entry{playerId, score, -1}
        heap.Push(this.entries, entry)
        this.playerIdx[playerId] = entry
    }
    // fmt.Println("addscore: ", this.String())
}


func (this *Leaderboard) Top(K int) int {
    sum := 0
    poped := []*Entry{}
    size := this.entries.Len()
    for i := 0; i < K; i++ {
        if i >= size {
            break
        }
        entryAny := heap.Pop(this.entries)
        entry := entryAny.(*Entry)
        //fmt.Println("poped: ", *entry)
        poped = append(poped, entry)
        sum += entry.Score
    }

    for _, entry := range poped {
            heap.Push(this.entries, entry)
    }
    return sum
}


func (this *Leaderboard) Reset(playerId int)  {
    if entry, ok := this.playerIdx[playerId]; ok {
        this.entries.Update(entry, 0)
    }
        // fmt.Println("reset: ", playerId, this.String())

}

func (this *Leaderboard) String() string {
    result := ""
    list := []*Entry(*this.entries)
    for i, entry := range list {
        result += fmt.Sprintf("%v(%d) ", *entry, i)
    }
    return result
}

/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */
