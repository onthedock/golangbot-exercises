package main

func main() {
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	map2 := map1

	if map1 == map2 { // ERROR: maps can't be compared (unless they're compared to 'nil')
	}
}
