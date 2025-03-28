import "slices"
type Set struct {
    c map[string]bool
    idx int
}

func (this *Set) IsSubSet(another *Set) bool {
    if len(this.c) > len(another.c) {
        return false
    }

    for key := range this.c {
        if another.c[key] == false {
            return false
        }
    }

    return true
}

func peopleIndexes(favoriteCompanies [][]string) []int {
    sets := make([]*Set, len(favoriteCompanies))

    // appears = make([]map[string]bool{}, len(favoriteCompanies))

    for i := range favoriteCompanies {
        companies := favoriteCompanies[i]
        m := map[string]bool{}
        // appears[i] = m
        for j := range companies {
            m[companies[j]] = true
        }
        sets[i] = &Set {
            c: m,
            idx: i,
        }
    }

    slices.SortFunc(sets, func (a, b *Set) int {
        return len(a.c) - len(b.c)
    })
    ans := []int{}

    for i := 0; i < len(sets) - 1; i++ {
        isSubSet := false
        for j := i+1; j < len(sets); j++ {
            set1 := sets[i]
            set2 := sets[j]
            if set1.IsSubSet(set2) {
                isSubSet = true
                break
            }
        }
        if !isSubSet {
            ans = append(ans, sets[i].idx)
        }
    }

    ans = append(ans, sets[len(sets) - 1].idx)
    slices.Sort(ans)
    return ans

}
