Similar to python gives index, value
```go
package main  
  
import "fmt"  
  
func main() {  
   pow := []int{1, 2, 4, 8, 16, 32, 64}  
   for i, v := range pow {  
      fmt.Printf("%d %d\n", i, v)  
   }  
}
```

