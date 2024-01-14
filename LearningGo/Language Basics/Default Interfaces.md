![[Pasted image 20240114050113.png]]
https://go.dev/tour/methods/19
![[Pasted image 20240114055310.png]]
readers
```go
// signature
func Read(b []byte) (n int, err error)
```
images
```go
type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```
[errors.go](https://github.com/adiChoudhary/learningGo/blob/main/code/PracticingGo/somethingMoreAdvanced/defaultInterfaces/errors.go)
[stringers.go](https://github.com/adiChoudhary/learningGo/blob/main/code/PracticingGo/excercises/stringers.go)
[errors.go](https://github.com/adiChoudhary/learningGo/blob/main/code/PracticingGo/excercises/errors.go)
