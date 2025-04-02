func canCompleteCircuit(gas []int, cost []int) int {
    sumGas := 0
    for _, g := range gas {
        sumGas += g
    }

    sumCost := 0
    for _, c := range cost {
        sumCost += c
    }

    if sumGas < sumCost {
        return -1
    }

    tank := 0
    ans := 0
    for i := range gas {
        // 这种情况是因为题目里面说明只有1个唯一的解
        // 同时因为上面已经判断过一定有解，那么如果其中某一半的gas-cost是负数，那么另一半一定是正数
        tank += gas[i]
        tank -= cost[i]
        if tank < 0 {
            ans = i+1
            tank = 0
        }
        // 如果是有多个解，就需要遍历所有开头判断是否成立
    }

    return ans
}
