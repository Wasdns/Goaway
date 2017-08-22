// learngo by examples: maps
package main

import "fmt"

func main() {
	// create a route table in terms of a map
	routeMap := make(map[string]int)
	routeMap["10.0.0.1/32"] = 1
	routeMap["10.0.0.2/32"] = 2
	fmt.Println("route table:", routeMap)

	port1 := routeMap["10.0.0.1/32"]
	fmt.Println("port1 physical number:", port1)
	fmt.Println("scale of route table:", len(routeMap))

	// delete an entry
	delete(routeMap, "10.0.0.1/32")
	fmt.Println("route table:", routeMap)

	// identify whether the entry has been deleted
	_, prs := routeMap["10.0.0.1/32"]
	if !prs {
		fmt.Println("Entry has been deleted!")
	}

	// directly declare a map
	directRouteMap := map[string]int{"10.0.0.3/32" : 3, "10.0.0.4/32" : 4}
	fmt.Println("direct route map:", directRouteMap)
}