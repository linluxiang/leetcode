import "math"

func equal(a, b float64) bool {
    if math.IsInf(a, 1) && math.IsInf(b, 1) {
        return true
    }

	if math.Abs(a-b) < 0.000001 {
		return true
	}
	return false
}

func getKN(a, b []int) (k float64, n float64) {
    x1, y1, x2, y2 := a[0], a[1], b[0], b[1]
    if x1 == x2 && y1 == y2 {
        return 0, 0
    }

    if y1 == y2 {
        return 0, float64(y1)
    }

    if x1 == x2 {
        return math.Inf(1), math.Inf(1)
    }

    k = float64(y2 - y1) / float64(x2 - x1)
    return k, float64(y1) - k*float64(x1)
}

func checkStraightLine(coordinates [][]int) bool {
	result := true
	var k, n *float64

	// k = float64(coordinates[1][1]-coordinates[0][1]) / float64(coordinates[1][0]-coordinates[0][0])
	for i := range coordinates {
        if i + 1 < len(coordinates) {
            newk, newn := getKN(coordinates[i], coordinates[i+1])

            if k == nil {
                k, n = &newk, &newn
                continue
            } else {
                            // fmt.Println(*k, *n, newk, newn)

                if !equal(*k, newk) || !equal(*n, newn) {
                    return false
                }
            }
        }

	}

	return result
}
