```go
func greatest() func(int) int {  
   cur := -1  
   // This is an anonymous function with access to variable declared outside its scope also known as closure  
   return func(x int) int {  
      if x > cur {  
         cur = x  
         return x  
      } else {  
         return cur  
      }  
   }  
}
```

![[Pasted image 20231128010526.png]]
With passing function as values with closures can create the above magic.

https://www.calhoun.io/5-useful-ways-to-use-closures-in-go/
