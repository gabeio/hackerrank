package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var (
	vPython int
	vC      int
	vJava   int
)

func main() {
	var err error
	var s string
	var eol bool = false
	java := regexp.MustCompile(`(java|import\ (java).*;|System\.)`)
	c := regexp.MustCompile(`(\#include[\ ]?[\<\"]\w+(\.h)?[\>\"]|cin\ ?>>|cout\ ?<<|system\(|using namespace)`)
	python := regexp.MustCompile(`((\#\ .*)|[\^\ \t]print\ .*|[\^\ \t]print\(.*\)|\w+\[\d+\:\d+\])`)
	loop := true
	bio := bufio.NewReader(os.Stdin)
infinite:
	for loop && !eol {
		s, err = bio.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR| %s\n", err)
			loop = false
			break infinite
		}
		if len(s) > 0 {
			fmt.Fprintf(os.Stderr, "s|%s", s)
			switch {
			case python.MatchString(s):
				vPython += 1
				fmt.Fprintf(os.Stderr, "vote| Python\n")
			case java.MatchString(s):
				vJava += 1
				fmt.Fprintf(os.Stderr, "vote| Java\n")
			case c.MatchString(s):
				vC += 1
				fmt.Fprintf(os.Stderr, "vote| C\n")
			default:
				//loop
			}
		}
	}
	if vPython != 0 || vJava != 0 || vC != 0 {
		if (vPython > vJava) && (vPython > vC) {
			fmt.Print("Python")
		} else if vJava > vC {
			fmt.Print("Java")
		} else {
			fmt.Print("C")
		}
	} else {
		fmt.Fprintf(os.Stderr, "unknown language\n")
	}
}
