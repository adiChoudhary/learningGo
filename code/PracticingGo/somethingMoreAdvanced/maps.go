package main

import "fmt"

type Geometry struct {
	length, width, height int
}

func main() {
	var mp map[string]Geometry     // zero value of maps is nil and keys can't be added to it later if its not initialized
	mp = make(map[string]Geometry) // gives an initialized map
	mp["Cube"] = Geometry{3, 3, 3}
	mp["Cuboid"] = Geometry{1, 2, 3}
	fmt.Println(mp["Cube"])
	fmt.Println(mp)

	//Another way to initialize your map
	anotherMp := map[string]Geometry{
		"Cuboid": Geometry{1, 2, 3},
		"Cube":   Geometry{2, 2, 2},
	}
	anotherMp["NayaCuboid"] = Geometry{4, 5, 6}
	fmt.Println(anotherMp)

	//If top level just a type name it can be omitted
	anotherMpV2 := map[string]Geometry{
		"Cuboid": {1, 2, 3},
		"Cube":   {2, 2, 2},
	}
	anotherMpV2["NayaCuboid"] = Geometry{4, 5, 6}
	fmt.Println(anotherMpV2)

	// Basic operations on map
	// Update m[key] = elem
	anotherMpV3 := map[string]Geometry{
		"Cuboid": {1, 2, 3},
		"Cube":   {2, 2, 2},
	}
	fmt.Println(anotherMpV3)
	anotherMpV3["Cuboid"] = Geometry{length: 2, height: 5}
	fmt.Println(anotherMpV3)

	// Retrieve elem = m[key]
	fmt.Println(anotherMpV3["Cube"])

	// Remove elem delete(m, key)
	delete(anotherMpV3, "Cube")
	fmt.Println(anotherMpV3)

	// Test if key present
	elem, ok := anotherMpV3["Cuboid"]
	fmt.Printf("%v %v", elem, ok)
}
