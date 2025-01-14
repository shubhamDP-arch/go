// package main

// import "fmt"

// func main() {
//     f()
//     fmt.Println("Returned normally from f.")
// }

// func f() {
//     defer func() {
//         if r := recover(); r != nil {
//             fmt.Println("Recovered in f", r)
//         }
//     }()
//     fmt.Println("Calling g.")
//     g(0)
//     fmt.Println("Returned normally from g.")
// }

// func g(i int) {
//     if i > 3 {
//         fmt.Println("Panicking!")
//         panic(fmt.Sprintf("%v", i))
//     }
//     defer fmt.Println("Defer in g", i)
//     fmt.Println("Printing in g", i)
//     g(i + 1)
// }

// package main

// import (
// 	"context"
// 	"fmt"
// )

// func enrichContext(ctx context.Context) context.Context {
//     return context.WithValue(ctx, "request-id", "12345")
// }

// func main()  {
//     fmt.Println("Go context")
//     ctx := context.Background()
//     ctx = enrichContext(ctx)
//     doSomethingColl(ctx)    
// }

// package main

// import (
//     "fmt"
//     "time"
// )

// func worker(id int) {
//     fmt.Printf("Worker %d starting\n", id)
//     time.Sleep(2 * time.Second) // Simulate work
//     fmt.Printf("Worker %d done\n", id)
// }

// func main() {
//     for i := 1; i <= 3; i++ {
//         go worker(i) // Launch each worker as a goroutine
//     }

//     time.Sleep(3 * time.Second) // Wait for goroutines to finish
//     fmt.Println("Main function finished")
// }

// package main

// import (
//     "fmt"
//     "sync"
// )

// func worker(id int, jobs <-chan int, results chan<- int) {
//     for job := range jobs {
//         fmt.Printf("Worker %d processing job %d\n", id, job)
//         results <- job * 2 // Send result back
//     }
// }

// func main() {
//     const numWorkers = 3
//     const numJobs = 5

//     jobs := make(chan int, numJobs)
//     results := make(chan int, numJobs)

//     var wg sync.WaitGroup

//     for w := 1; w <= numWorkers; w++ {
//         wg.Add(1)
//         go func(id int) {
//             defer wg.Done()
//             worker(id, jobs, results)
//         }(w)
//     }

//     for j := 1; j <= numJobs; j++ {
//         jobs <- j
//     }
//     close(jobs)

//     wg.Wait()
//     close(results)

//     for result := range results {
//         fmt.Println("Result:", result)
//     }
// }

package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}