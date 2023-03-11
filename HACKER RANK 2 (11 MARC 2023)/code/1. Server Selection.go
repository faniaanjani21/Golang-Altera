package main

import (
    "fmt"
    "sort"
)

func getMinServers(expectedLoad int, servers []int) int {
    n := len(servers)
    sort.Ints(servers)

    count := 0
    sum := 0
    for reza := n - 1; reza >= 0; reza-- {
        sum += servers[reza]
        count++
        if sum >= expectedLoad {
            return count
        }
    }

    return -1
}

func main() {
    var expectedLoad, n int
    fmt.Scan(&expectedLoad)

    fmt.Scan(&n)

    servers := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&servers[i])
    }

    minServers := getMinServers(expectedLoad, servers)
    if minServers == -1 {
        fmt.Println(-1)
    } else {
        fmt.Printf("%d\n", minServers)
    }
}