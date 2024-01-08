![[Pasted image 20231209154925.png]]
Each value(day, hours, minutes etc.) are assigned a predetermined value in Go's layout using which it determines how to replace those values in string with the correct value.

Using that we can create the above output

```go
package main  
  
import (  
   "fmt"  
   "time")  
  
func takeInput(data interface{}) {  
   _, err := fmt.Scan(data)  
   if err != nil {  
      fmt.Printf("Error occured %v", err)  
   }  
}  
  
func solve(currTime time.Time, offset int) time.Time {  
   return currTime.Add(time.Duration(int(time.Hour) * offset))  
}  
  
func main() {  
   value := time.Now().UTC()  
   var offset int  
   fmt.Print("Enter the time zone offset to GMT:")  
   takeInput(&offset)  
   value = solve(value, offset)  
   // layout is according to predefined RFCs, use them to define which value is needed where  
   fmt.Println("The current time is:", value.Format("15💀04💀05😁"))  
}
```

For referring to layouts 
https://yourbasic.org/golang/format-parse-string-time-date-example/#layout-options