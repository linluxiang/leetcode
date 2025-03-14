import "fmt"

type Leaderboard struct {
    data map[int]int // player id to score
}


func Constructor() Leaderboard {
    return Leaderboard{
        data: map[int]int{},
    }
}
OAOAOAOAOA

func (this *Leaderboard) AddScore(playerId int, score int)  {
    this.data[playerId] += score
}


func (this *Leaderboard) Top(K int) int {
    l := []int{}
    for _, v := range this.data {
        l = append(l, v)
    } 
    return quickSelect(l, K, 0, len(l) - 1)
}

func quickSelect(l []int, k int, begin, end int) int {
    sum := 0
    if len(l) == k {
        for _, v := range l {
            sum += v
        }
        return sum
    }
    if k == 0 {
        return 0
    }


    smaller := []int{}
    equal := []int{}
    bigger := []int{}
    // fmt.Println(k, begin, end, l)
    pivot := l[k-1]
    for i := begin; i <= end; i++ {
        if l[i] < pivot {
            smaller = append(smaller, l[i])
        }
        if l[i] == pivot {
            equal = append(equal, l[i])
        }
        if l[i] > pivot {
            bigger = append(bigger, l[i])
        }
    }
    // fmt.Println(k, bigger, equal, smaller)


    if len(bigger) + len(equal) < k {
        for i := 0; i < len(bigger); i++ {
            sum += bigger[i]
        }
        for i := 0; i < len(equal); i++ {
            sum += equal[i]
        }
        sum += quickSelect(smaller, k-len(bigger) - len(equal), 0, len(smaller) - 1)
    } else if len(bigger) <= k {
        for i := 0; i < len(bigger); i++ {
            sum += bigger[i]
        }
        for i := 0; i < k - len(bigger); i++ {
            sum += equal[i]
        }
        return sum
    } else { //len(bigger) > k
        sum += quickSelect(bigger, k, 0, len(bigger) - 1)
        return sum
    }

    return sum
}

func (this *Leaderboard) Reset(playerId int)  {
    this.data[playerId] = 0
}


/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */
