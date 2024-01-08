https://gobyexample.com/string-formatting#:~:text=To%20left%2Djustify%2C%20use%20the%20%2D%20flag.&text=You%20may%20also%20want%20to,For%20basic%20right%2Djustified%20width.&text=To%20left%2Djustify%20use%20the%20%2D%20flag%20as%20with%20numbers.
![[Pasted image 20231212001807.png]]
```go
package main  
  
import (  
   "fmt"  
   "math")  
  
func main() {  
   for i := 0; i <= 5; i++ {  
      if i == 0 {  
         fmt.Printf("%-10s%-10s%-10s\n", "a", "b", "pow(a, b)")  
      } else {  
         fmt.Printf("%-10d%-10d%-10d\n", i, i+1, int(math.Pow(float64(i), float64(i+1))))  
      }  
   }  
}
```