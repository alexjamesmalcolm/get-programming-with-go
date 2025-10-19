package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %v° C.\n", temp)
	temperature["Earth"] = 16
	temperature["Venus"] = 464
	fmt.Println(temperature)

	moon := temperature["Moon"]
	fmt.Println(moon)

	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %v° C.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}

	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}
	planetsMarkII := planets
	planets["Earth"] = "whoops"
	fmt.Println(planets)
	fmt.Println(planetsMarkII) // Also has Earth's value updated, because maps are not copied.
	delete(planets, "Earth")
	fmt.Println(planetsMarkII) // Remove Earth from the map
}
