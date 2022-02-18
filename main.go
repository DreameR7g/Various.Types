package main

import (
	"fmt"
	"time"
)

type Character struct {
	// name
	Name string
	// job
	Job string
	// level
	Level int
	// race
	Race string
	// Current Date and Time
	Date time.Time
}

func main() {
	fmt.Println("Light Party")
	fmt.Println(Character{
		Name:  "Dreamer Rigorstorm",
		Job:   "Monk",
		Level: 90,
		Race:  "Lalafell",
		Date:  time.Now(),
	})
	fmt.Println(Character{
		Name:  "Siryu Unslay",
		Job:   "Dark Knight",
		Level: 90,
		Race:  "Au'Ra",
		Date:  time.Now(),
	})
	fmt.Println(Character{
		Name:  "James Carmike",
		Job:   "Sage",
		Level: 90,
		Race:  "Viera",
		Date:  time.Now(),
	})
	fmt.Println(Character{
		Name:  "Kisa Rozen",
		Job:   "Reaper",
		Level: 90,
		Race:  "Hyur",
		Date:  time.Now(),
	})
}
