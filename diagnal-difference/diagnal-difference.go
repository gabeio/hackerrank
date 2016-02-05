package main
import (
    "fmt"
    "math"
)

func main() {
    var square,left,right int
    fmt.Scanf("%d", &square)
    val := make([]int, square*square)
    for i:=0; i<square*square; i++ {
        fmt.Scanf("%d", &val[i])
    }
    for i:=0; i<square*square;i+=(square+1) {
        left += val[i]
    }
    for i:=(square-1); i<((square*square)-1);i+=(square-1) {
        right += val[i]
    }
    fmt.Printf("%d", int(math.Abs(float64(left-right))))
}
