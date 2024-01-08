```go
package main  
  
import (  
   "fmt"  
   "strconv")  
  
func main() {  
   sum := 0  
   for i := 0; i < 10; i++ {  
      sum += i  
   }  
   fmt.Println(sum)  
  
   // For is Go's while  
   temp := 1  
   for temp < 20 {  
      temp++  
      fmt.Println("Hello World" + string(rune(65)))  
   }  
  
   // if you omit conditional, we get simple infinite loop😁  
   temp = 1  
   for {  
      if temp == 20 {  
         break  
      }  
      fmt.Println("Hello World" + strconv.Itoa(temp))  
      temp++  
   }  
}
```
