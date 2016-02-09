package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type (
	Stack []byte
)

func (s *Stack) Push(i byte) {
	(*s) = append((*s), i)
}

func (s *Stack) Pop() byte {
	// var i byte = (*s)[len(*s)-1]
	// (*s) = (*s)[:len(*s)-1]
	// return i
	return (*s).Remove((len(*s) - 1))
}

func (s *Stack) Remove(i int) byte {
	if len(*s) > 0 {
		var b byte = (*s)[i]
		var right []byte
		if (len(*s) - 1) > i {
			right = (*s)[i+1:]
		}
		(*s) = (*s)[:i]
		(*s) = append((*s), right...)
		return b
	} else {
		return 0x00
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	bio := bufio.NewReader(os.Stdin)
	// top:
	for i := 0; i < n; i++ { // for lines
		var st Stack = Stack{}
		var tf bool = true
		var cl bool = false
	inf:
		for {
			fmt.Fprintf(os.Stderr, "st| %+s\n", st)
			b, err := bio.ReadByte()
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR| %s\n", err)
				break inf
			} else if b == 0x0A { // nl break
				cl = true
				break inf
			}
			if bytes.Equal([]byte{b}, []byte{'{'}) || bytes.Equal([]byte{b}, []byte{'('}) || bytes.Equal([]byte{b}, []byte{'['}) {
				fmt.Fprintf(os.Stderr, "push| %s\n", string(b))
				st.Push(b)
			} else if bytes.Equal([]byte{b}, []byte{')'}) {
				z := st.Pop()
				fmt.Fprintf(os.Stderr, "pop| %s\n", string(b))
				if !bytes.Equal([]byte{b}, []byte{z + 1}) {
					tf = false
					break inf
				}
			} else if bytes.Equal([]byte{b}, []byte{'}'}) || bytes.Equal([]byte{b}, []byte{']'}) {
				z := st.Pop()
				if z == 0x00 {
					tf = false
					break inf
				}
				fmt.Fprintf(os.Stderr, "epop| %s\n", string(b))
				if !bytes.Equal([]byte{b}, []byte{z + 2}) {
					tf = false
					break inf
				}
			} else {
				tf = false
				break inf
			}
		}
		if len(st) > 0 { // if there are any left over
			tf = false // false
		}
		switch tf {
		case true:
			fmt.Fprintf(os.Stdout, "YES\n")
		case false:
			if cl == false { // clear line only if it wasn't already cleared
				_, _ = bio.ReadString('\n') // clear line
			}
			fmt.Fprintf(os.Stdout, "NO\n")
		}
	}
}
