package main

import "fmt"

func main() {

	myMap := make(map[string]int)
	myMap["key1"] = 1
	myMap["key2"] = 10
	myMap["key3"] = 10
	myMap["key4"] = 10
	myMap["code"] = 10
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])
	myMap["code"] = 21
	fmt.Println(myMap)
	// delete(myMap, "key1")
	fmt.Println(myMap)
	// clear(myMap)
	fmt.Println(myMap)
	// myMap["key1"] = 100
	_, ok := myMap["key1"]
	if ok {
		fmt.Println("Is a value ascociated with key1: ", ok)
	} else {
		fmt.Println("The Key1 is not present...")
	}

	myMap2 := map[string]int{"a": 1, "b": 2}
	for key, value := range myMap2 {
		fmt.Println("key: ", key, "Value: ", value)
	}
	myMap3 := make(map[string]map[string]int)
	myMap3["map1"] = myMap2

}
