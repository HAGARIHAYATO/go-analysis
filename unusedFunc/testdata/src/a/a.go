package a

import "fmt"

func f() {
	fmt.Print("unused")
}

func G() {
	fmt.Println("used")
	e()
}

func e() {
	fmt.Println("too")
}

