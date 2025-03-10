import "slices"
import "strings"
import "fmt"

type Itineraries struct {
    data []string
}

func NewItineraries() *Itineraries {
    return &Itineraries {
        data: []string{},
    }
}

func (i *Itineraries) Add(dest string) {
    pos, _ := slices.BinarySearchFunc(i.data, dest, func (a, b string) int {
        return strings.Compare(a, b)
    })

    i.data = slices.Insert(i.data, pos, dest)
}

func (i *Itineraries) Pop() string {
    if len(i.data) == 0 {
        return ""
    }
    r := i.data[0]
    i.data = i.data[1:]
    return r
}

// 这里是一个欧拉图，也就是说每一条边必须访问一次
// 用DFS比BFS更容易。

func traverse(its map[string]*Itineraries, city string, result *[]string) {
    
    it := its[city]

    if it == nil {
        // found end
        *result = append(*result, city)
        return
    }

    dest := it.Pop() // 用Pop删除已经访问过的边
    for dest != "" {
        traverse(its, dest, result)
        dest = it.Pop()
    }
    
    *result = append(*result, city)
}

func findItinerary(tickets [][]string) []string {
    its := map[string]*Itineraries{}
    for _, t := range tickets {
        from := t[0]
        to := t[1]
        if its[from] == nil {
            its[from] = NewItineraries()
        }
        its[from].Add(to)
    }

    result := []string{}
    traverse(its, "JFK", &result)

    slices.Reverse(result) // DFS本质是一个栈，因为不是在遍历过程中输出，所以需要倒序

    return result
}
