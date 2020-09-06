package a

import "fmt"

func f() {
	fmt.Print("f")
}

func G() {
	fmt.Println("G")
}

func aaa() {
	_ = func() { fmt.Println("no implement function") }
}
