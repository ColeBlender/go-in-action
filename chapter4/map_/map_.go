package map_

import "fmt"

func Map_() {
	dict := make(map[string]int)
	fmt.Println(dict)

	d1 := map[string]int{"cool": 1, "dank": 2}

	fmt.Println(d1)

	delete(d1, "cool")

	fmt.Println(d1)

	fmt.Println(len(d1))

}
