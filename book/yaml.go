// package main

// import(
// 	"fmt"
// 	// temp "github.com/go-yaml/yaml"
// )

// func main()  {
// 	fmt.Println("Hello, 世界")
// 	output, _ := temp.Marshal(map[string]string{"a": "b", "c": "d"})
// 	fmt.Println(string(output))
// }
// package main
// import "fmt"
// func main() {
// l := TroubleMessagePrinter{}
// PrintSomething(l)
// }
// func PrintSomething(m messagePrinter) {
// m.Print()
// }
// type messagePrinter interface {
// Print()
// }
// type LoveMessagePrinter struct{}
// func (l LoveMessagePrinter) Print() {
// fmt.Println("I love Golang")
// }

// type TroubleMessagePrinter struct{}
// func (t TroubleMessagePrinter) Print() {
// fmt.Println("I am still confused by this")
// }
// func (t TroubleMessagePrinter) AdditionalFunc() {
// fmt.Println("additionalFunc")
// }

// package main
// import "fmt"
// func main() {
//  allTasks := []Subtask{Subtask{Status: "incomplete"}, Subtask{Status:
// "completed"}}
//  for _, x := range allTasks {
// 	fmt.Println(x.Status)
//  if x.Status != "completed" {
// 	fmt.Println("Main task is still incomplete")
//  }
//  }
// }
// type Subtask struct {
// Param1 string
// Param2 string
// Status string
// }

// func main() {
// 	items := map[string]string{
// 	 "key1": "value1",
// 	 "key2": "value2",
// }
// for idx, x := range items {
// 	fmt.Println(idx)
// 	fmt.Println(x)
//  }
//  }

// package main
// import (
// "fmt"
// "time"
// )
// func main() {
// fmt.Println("Start main")
// go side()
// fmt.Println("Return to main")
// time.Sleep(5 * time.Second)
// fmt.Println("End main")
// }
// func side() {
// fmt.Println("Start side process")
// time.Sleep(3 * time.Second)
// fmt.Println("End side process")
// }
// 

package main
import (
 "fmt"
 "log"
 "net/http"
)
func handler(w http.ResponseWriter, r *http.Request) {
 log.Print("Hello world received a request.")
 defer log.Print("End hello world request")
 fmt.Fprintf(w, "Hello World")
}
func main() {
 log.Print("Hello world sample started.")
 http.HandleFunc("/", handler)
 http.ListenAndServe(":8080", nil)
}