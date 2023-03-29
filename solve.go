package main

import (
    "fmt"
    "math"
)

type Coord struct {
    X, Y float64
}

func dist(p1, p2 Coord) float64 {
    dx := p1.X - p2.X
    dy := p1.Y - p2.Y
    return math.Sqrt(dx*dx + dy*dy)
}

func main() {
    var n int
    fmt.Scan(&n)

    coords := make([]Coord, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&coords[i].X, &coords[i].Y)
    }

    path := make([]int, n)
    for i := 0; i < n; i++ {
        path[i] = i
    }

    res := 1e9
    for {
        d := 0.0
        for i := 0; i < n-1; i++ {
            d += dist(coords[path[i]], coords[path[i+1]])
        }
        if d < res {
            res = d
        }
        if !nextPermutation(path) {
            break
        }
    }

    fmt.Printf("%v\n", res)
}

func nextPermutation(a []int) bool {
    i := len(a) - 2
    for ; i >= 0 && a[i] >= a[i+1]; i-- {}
    if i < 0 {
        return false
    }
    j := len(a) - 1
    for ; a[j] <= a[i]; j-- {}
    a[i], a[j] = a[j], a[i]
    for k, m := i+1, len(a)-1; k < m; k, m = k+1, m-1 {
        a[k], a[m] = a[m], a[k]
    }
    return true
}
