package main
import (
	"fmt"
	"time"
)

func sub() {
	for {
		fmt.Println("sub loop")
		time.Sleep(100 * time.Millisecond)
	}
}

// goをつけないと、sub()が永久ループし、main loopが表示されない
// sub loopとmain loopが交互に表示される
func main() {
	go sub()

	for {
		fmt.Println("main loop")
		time.Sleep(200 * time.Millisecond)
	}
}
