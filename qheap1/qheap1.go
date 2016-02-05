package main
import "fmt"

type (
	heap struct {
		ints []int64
		small int64
	}
)

func (h *heap) Push(a int64) {
	check := h.small
	if a < check || len(h.ints)==0 {
		h.small = a
	}
	h.ints = append(h.ints, a)
}

func (h *heap) Pop(a int64) {
	find := false
	found := false
	old := h.ints
	n := len(old)
	new := old[0:n-1]
	if a == h.small {
		find = true
	}
	for i := 0; i<n-1; i++ {
		if old[i] == a {
			found = true
		}
		if found {
			new[i] = old[i+1] // if found skip and continue
		}else{
			new[i] = old[i]
		}
	}
	if find == true && len(new)>0 {
		var min int64 = new[0]
		for i := 1; i<len(new); i++ {
			if new[i] < min {
				min = new[i]
			}
		}
		h.small = min
	}
	h.ints = new[:]
}

func (h *heap) Least() int64 {
	return h.small
}

func main() {
    var n,z int64
    array := &heap{}
    fmt.Scanf("%d", &n) // transactions
    for z = 0; z < n; z++ {
        var control int64
        fmt.Scanf("%d", &control)
        switch control {
        case 1: // add
            var add int64
            fmt.Scanf("%d",&add)
            array.Push(add)
        case 2: // delete
            var del int64
            fmt.Scanf("%d",&del)
            array.Pop(del)
        case 3: // print64 min
            fmt.Printf("%d\n", array.Least())
        }
    }
}
