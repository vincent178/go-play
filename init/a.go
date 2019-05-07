package init

import "fmt"

func A() {
	fmt.Println("a")
}

func init() {
	fmt.Println("init ch2 a")
}
