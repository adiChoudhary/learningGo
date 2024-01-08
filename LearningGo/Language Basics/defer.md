```go
package main  
  
import "fmt"  
  
func main() {  
   fmt.Println("1")  
  
   // all deferred function calls will be after the standard function calls  
   // all deferred functions will execute in LIFO fashion   defer fmt.Println("5")  
   defer fmt.Println("4")  
   fmt.Println("2")  
   fmt.Println("3")  
}
```

A go functionality to deprioritize a function call