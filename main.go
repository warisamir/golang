package main
import(
	"fmt"package main
	
	import (
		"fmt"
		"os"
		"time"
	)
	
	func main() {
		fmt.Println("Shortcut triggered! 🟢")
		f, _ := os.OpenFile("shortcut_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		f.WriteString(fmt.Sprintf("Triggered at %v\n", time.Now()))
	}
)
func main(){
	fmt.Println("hello world")
}