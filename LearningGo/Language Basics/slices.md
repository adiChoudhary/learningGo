provides a flexible view into the elements of array
![[Pasted image 20231126005641.png]]
![[Pasted image 20231126005757.png]]

How length and capacity in slices is handled
![[Pasted image 20231126012021.png]]
![[Pasted image 20240107180131.png]]
In the above highlighted line, capacity changes cause initial point changes
#### Appending to a slice
```go
var incSlice []int  
// It works on nil slices also  
incSlice = append(incSlice, 1)  
printSlice(incSlice)  
// it grows, parent array behind it gets replaced with a new one if its capacity is smaller  
incSlice = append(incSlice, 2)  
printSlice(incSlice)  
// We can add multiple elements at a time  
incSlice = append(incSlice, 3, 4, 5, 6, 7, 8)  
printSlice(incSlice)  
// We can also unpack a slice to append its elements to another slice  
incSlice = append(incSlice, []int{10, 20, 30}...)  
someNewSlice := make([]int, 6)  
incSlice = append(incSlice, someNewSlice...)  
printSlice(incSlice)
```

Creating Matrics in Go using slices(So Easy :))
```go
func Pic(dx, dy int) [][]uint8 {  
   picSlice := make([][]uint8, dy)  
   for i := 0; i < len(picSlice); i++ {  
      picSlice[i] = make([]uint8, dx)  
   }  
   return picSlice  
}
```