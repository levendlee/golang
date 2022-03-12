// https://go.dev/tour/concurrency/8

package main

import "fmt"
import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if (t.Left != nil) {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if (t.Right != nil) {
		Walk(t.Right, ch)
	}
}

func WalkAndClose(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

func CheckMatch(ch1 chan int, ch2 chan int) bool {
	for {
		var v1, v2 int
		var ok1, ok2 bool
		v1, ok1 = <- ch1
		v2, ok2 = <- ch2
		fmt.Println("Return values: v1=%s ok1=%s v2=%s ok2=%s", v1, ok1, v2, ok2)
		if (ok1 && ok2) {
			if (v1 != v2) {
				return false
			}
		}
		if (!ok1 && !ok2) {
			return true
		}
		if (!ok1 || !ok2) {
			return false
		}
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go WalkAndClose(t1, ch1)
	go WalkAndClose(t2, ch2)
	return CheckMatch(ch1, ch2)
}

func main() {
	ret := Same(tree.New(10), tree.New(10))
	fmt.Println(ret)
}
