
import (
	"fmt"
    "container/heap"
)


type Queue []int

func (q *Queue) Push(v any) {
    entry := v.(int)

    *q = append(*q, entry)
}

func (q *Queue) Pop() any {
    old := *q
    size := len(old)
    if size == 0 {
        return nil
    }
    last := old[size - 1]
    *q = old[:size - 1]
    return last
}

func (q Queue) Len() int {
    return len(q)
}

func (q Queue) Less(a, b int) bool {
    return q[b] < q[a]
}

func (q Queue) Swap(i, j int) {
    q[i], q[j] = q[j], q[i]
}


type Leaderboard struct {
    playerScore map[int]int // id to score
}


func Constructor() Leaderboard {
    return Leaderboard{
        playerScore: map[int]int{},
    }
}


func (this *Leaderboard) AddScore(playerId int, score int)  {
    this.playerScore[playerId] += score
    // fmt.Println("addscore: ", this.String())
}


func (this *Leaderboard) Top(K int) int {
    sum := 0
    q := make(Queue, len(this.playerScore))
    i := 0
    for _, score := range this.playerScore {
        q[i] = score
        i++
    }

    heap.Init(&q)

    for range K {
        v := heap.Pop(&q)
        sum += v.(int)
    }

    return sum
}


func (this *Leaderboard) Reset(playerId int)  {
    this.playerScore[playerId] = 0
}


/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */
