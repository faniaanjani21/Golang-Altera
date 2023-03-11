package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func rearrange(arr []int) []int {
    n := len(arr)

    // Sort the array in ascending order
    sort.Ints(arr)

    // Rearrange the array by alternating the smallest and largest elements
    for i := 0; i < n/2; i++ {
        j := n - 1 - i*2
        arr[i], arr[j] = arr[j], arr[i]
    }

    return arr
}

func main() {
    // Read input from stdin
    reader := bufio.NewReader(os.Stdin)

    // Read the size of the array
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    n, _ := strconv.Atoi(input)

    // Read the elements of the array
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        arr[i], _ = strconv.Atoi(input)
    }

    // Rearrange the array
    arr = rearrange(arr)

    // Calculate U
    u := float64(arr[0])
    for i := 1; i < n; i++ {
        if i%2 == 1 {
            u *= float64(arr[i])
        } else {
            u *= float64(arr[i]) / float64(arr[i-1])
        }
    }

    // Print the rearranged array
    for _, x := range arr {
        fmt.Println(x)
    }

    // Print the value of U
    fmt.Println(u)
}