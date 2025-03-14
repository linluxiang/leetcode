func dfs(graph [][]int, id int, path []int, result *[][]int) {
    path = append(path, id)

    if id == len(graph) - 1 {
        // 这里copy很重要，因为path其实是一个胖指针，如果直接存胖指针，那么有可能该指针指向的内容会在后面被修改
        newPath := make([]int, len(path))
        copy(newPath, path)
        *result = append(*result, newPath)
    }
    for _, v := range graph[id] {
        dfs(graph, v, path, result)
    }
}

func allPathsSourceTarget(graph [][]int) [][]int {
    result := [][]int{}
    dfs(graph, 0, []int{}, &result)
    return result
}
