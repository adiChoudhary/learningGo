```go
func takeInput(data interface{}) {  
   _, err := fmt.Scan(data)  
   if err != nil {  
      fmt.Printf("Error occured %v", err)  
   }  
}
```

