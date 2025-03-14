// import "slices"

type Person struct {
    Id int
    like map[int]int // player id to idx
}

func NewPerson(id int, preference []int) *Person {
    like := make(map[int]int, len(preference))
    for i, playerId := range preference {
        like[playerId] = i
    }
    return &Person{
        Id: id,
        like: like,
    }
}

func (p *Person) LikeMore(a, b int) bool {
    return p.like[a] < p.like[b]
}


type Pair struct {
    p1 *Person
    p2 *Person
}

func NewPair(ids []int, people map[int]*Person) *Pair {
    return &Pair{people[ids[0]], people[ids[1]]}
}

func (p Pair) Another(id int) *Person {
    if p.p1.Id == id {
        return p.p2
    }
    return p.p1
}

func isUnhappy(person int, id2Pair map[int]*Pair, preferences [][]int, people map[int]*Person) bool {
    pair := id2Pair[person]
    another := pair.Another(person)
    r := false
    for _, morelikeId := range preferences[person] {
        if morelikeId == another.Id {
            break
        }
        pair2 := id2Pair[morelikeId]
        morelikePerson := people[morelikeId]
        anotherInPair2 := pair2.Another(morelikeId)
        if morelikePerson.LikeMore(person, anotherInPair2.Id) {
            return true
        }

    }
    return r
}

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
    people := make(map[int]*Person, n)

    for i, pref := range preferences {
        people[i] = NewPerson(i, pref)
    }

    ps := make([]*Pair, len(pairs))

    id2Pair := map[int]*Pair{}

    for i, pair := range pairs {
        ps[i] = NewPair(pair, people)
        id2Pair[pair[0]] = ps[i]
        id2Pair[pair[1]] = ps[i]
    }

    unhappy := 0

    for _, pair := range pairs {
        // pr := ps[i]
        for _, personId := range pair {
            if isUnhappy(personId, id2Pair, preferences, people) {
                unhappy += 1
            }
        }
    }

    return unhappy
}
