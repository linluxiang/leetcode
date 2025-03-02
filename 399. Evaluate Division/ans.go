func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    varSets := map[string]*map[string]float64 {}

    for i, eq := range equations {
        left := varSets[eq[0]]
        right := varSets[eq[1]]

        if left == nil && right == nil {
            set := map[string]float64{}
            set[eq[0]] = 1.0
            set[eq[1]] = 1.0 / values[i]
            varSets[eq[0]] = &set
            varSets[eq[1]] = &set
            continue
        }

        if left == nil && right != nil {
            set := *right
            set[eq[0]] = values[i]*set[eq[1]]
            varSets[eq[0]] = right
            continue
        }

        if left != nil && right == nil{
            set := *left
            set[eq[1]] = set[eq[0]] / values[i]
            varSets[eq[1]] = left
            continue
        }

        // left !=nil && right != nil
        // merge right to left
        setLeft := *varSets[eq[0]]
        newB := setLeft[eq[0]] / values[i]
        setLeft[eq[1]] = newB
        oldSetRight := *varSets[eq[1]]
        oldB := oldSetRight[eq[1]]
        for k, v := range oldSetRight {
            setLeft[k] = v * newB / oldB
            varSets[k] = varSets[eq[0]] 
        }
    }

    result := make([]float64, len(queries))
    for i, q := range queries {
        if varSets[q[0]] == nil || varSets[q[1]] == nil {
            result[i] = -1.0
            continue
        }

        if varSets[q[0]] != varSets[q[1]] {
            result[i] = -1.0
            continue
        }

        set := *varSets[q[0]]
        result[i] = set[q[0]] / set[q[1]]
    }

    return result
}
