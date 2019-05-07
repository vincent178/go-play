package init

import "fmt"

func B() {
	fmt.Println("b")
}

func init() {
	fmt.Println("init ch2 b")
}
