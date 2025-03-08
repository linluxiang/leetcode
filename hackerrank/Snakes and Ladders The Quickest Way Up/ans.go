package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'quickestWayUp' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY ladders
 *  2. 2D_INTEGER_ARRAY snakes
 */
 
type Node struct {
    Val int32
    Step int32
}

func quickestWayUp(ladders [][]int32, snakes [][]int32) int32 {
    // Write your code here
    smap := map[int32]int32{}
    for _, v := range snakes {
        smap[v[0]] = v[1]
    }
    
    lmap := map[int32]int32{} // from the dest to the start
    for _, v := range ladders {
        lmap[v[0]] = v[1]
    }
    
    step := map[int32]int32{}
    step[1] = 0
    step[2] = 1
    step[3] = 1
    step[4] = 1
    step[5] = 1
    step[6] = 1
    step[7] = 1
    
    queue := []*Node{
        {1, 0},
    }
    
    for len(queue) != 0 {
        head := queue[0]
        queue = queue[1:]
        // fmt.Println("visit: ", head)
        
        if v, ok := step[head.Val]; ok {
            // already visited
            // expand only when the Step is less than the cached value
            if v < head.Step {
                continue
            }
        }
        step[head.Val] = head.Step
        
        if v := lmap[head.Val]; v != 0 {
            queue = append(queue, &Node{v, head.Step})
            continue
        }
        
        if v := smap[head.Val]; v != 0 {
            queue = append(queue, &Node{v, head.Step})
            continue
        }
    
        
        for i := int32(1); i <= 6; i++ {
            if head.Val + i <= 100 {
                queue = append(queue, &Node{head.Val + i, head.Step + 1})
            }
        }

    }
    
    if step[100] == 0 {
        return -1
    }
    return step[100]
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n := int32(nTemp)

        var ladders [][]int32
        for i := 0; i < int(n); i++ {
            laddersRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

            var laddersRow []int32
            for _, laddersRowItem := range laddersRowTemp {
                laddersItemTemp, err := strconv.ParseInt(laddersRowItem, 10, 64)
                checkError(err)
                laddersItem := int32(laddersItemTemp)
                laddersRow = append(laddersRow, laddersItem)
            }

            if len(laddersRow) != 2 {
                panic("Bad input")
            }

            ladders = append(ladders, laddersRow)
        }

        mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        m := int32(mTemp)

        var snakes [][]int32
        for i := 0; i < int(m); i++ {
            snakesRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

            var snakesRow []int32
            for _, snakesRowItem := range snakesRowTemp {
                snakesItemTemp, err := strconv.ParseInt(snakesRowItem, 10, 64)
                checkError(err)
                snakesItem := int32(snakesItemTemp)
                snakesRow = append(snakesRow, snakesItem)
            }

            if len(snakesRow) != 2 {
                panic("Bad input")
            }

            snakes = append(snakes, snakesRow)
        }

        result := quickestWayUp(ladders, snakes)

        fmt.Fprintf(writer, "%d\n", result)
    }

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

