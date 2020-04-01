package add

import "fmt"

func init() {
	fmt.Println("add包 init")
}
func Add(x, y int) int {
	return (x + y)
}
