![[Pasted image 20240114044011.png]]
Pros
This decoupling allows for any method(or methods to be attached to a given struct). 

Cons
But reduces our ability to identify how many methods are associated with a given type

![[Pasted image 20240114044916.png]]
Something like temp.M() where temp is null will give error in Java but in Go, a nil receiver will be invoked where invoked function would handle the nil value.
This provides us with us clean ways to handle errors

![[Pasted image 20240114045215.png]]
This allows us to deal with unknown types and use [[Type Assertion]] to determine the type at a later stage and work on it

Go provides us with some common interfaces.
[[Default Interfaces]]