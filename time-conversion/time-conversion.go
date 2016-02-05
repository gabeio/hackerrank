package main
import(
    "fmt"
)

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    var h,m,s int
    var pm string
    fmt.Scanf("%2d:%2d:%2d%2s",&h,&m,&s,&pm)
    if h==12 {
        h = 0
    }
    switch pm {
    case "AM":
        fmt.Printf("%02d:%02d:%02d",h,m,s)
    case "PM":
        fmt.Printf("%02d:%02d:%02d",(h+12),m,s)
    default:
        fmt.Println("unknown time format")
    }
}
