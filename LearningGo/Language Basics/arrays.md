In go, we can change the array size
Syntax
```go
package main  
  
import "fmt"  
  
func main() {  
   var list [2]int  
   list[0] = 1  
   list[1] = 2  
   fmt.Println(list)  
   strList := [2]string{"Hello", "World"}  
   fmt.Println(strList)  
}
```