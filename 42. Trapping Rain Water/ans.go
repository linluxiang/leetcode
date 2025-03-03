import "fmt"

func volumn(height []int, i, j int) int {
    // make sure that height[i] <= height[j]
    r := 0
    start := min(i, j)
    end := max(i, j)
    for p := start+1; p < end; p++ {
        r += height[i] - height[p]
    }
    // fmt.Println("v: ", i, j, r)

    return r
}

func trap(height []int) int {
    if len(height) == 1 {
        return 0
    }

    // scan from left to right
    barLeft := [][]int{}
    for i:=0;i<len(height);i++ {
        if height[i] == 0 {
            continue
        }
        if len(barLeft) == 0 && height[i] > 0 {
            barLeft = append(barLeft, []int{i, height[i]} )
            continue
        }

        if height[i] >= barLeft[len(barLeft) - 1][1] {
            barLeft = append(barLeft, []int{i, height[i]} )
        }
    }
    // fmt.Println(barLeft)
    barRight := [][]int{}
    lastBarLeftIdx := 0
    if len(barLeft) > 0 {
        lastBarLeftIdx = barLeft[len(barLeft) - 1][0]
    }
    for i:=len(height) - 1; i>=lastBarLeftIdx; i-- {
        if height[i] == 0 {
            continue
        }
        if len(barRight) == 0 && height[i] > 0 {
            barRight = append(barRight, []int{i, height[i]} )
            continue
        }

        if height[i] >= barRight[len(barRight) - 1][1] {
            barRight = append(barRight, []int{i, height[i]} )
        }
    }
    // fmt.Println(barRight)

    result := 0
    if len(barLeft) >= 2 {
        i, j := 0, 1
        for j < len(barLeft) {
            result += volumn(height, barLeft[i][0], barLeft[j][0])
            i++
            j++
        }
    }
    if len(barRight) >= 2 {
        i, j := 0, 1
        for j < len(barRight) {
            result += volumn(height, barRight[i][0], barRight[j][0])
            i++
            j++
        }
    }
    return result
}

