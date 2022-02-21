package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Character struct {
	// name
	Name string `json:"name"`
	// job
	Job string `json:"job"`
	// level
	Level int `json:"level"`
	// race
	Race string `json:"race"`
	// Current Date and Time
	Date time.Time
}

func main() {
	party := make(map[string]Character)

	fmt.Println("Light Party")
	party["Dreamer"] = Character{
		Name:  "Dreamer Rigorstorm",
		Job:   "Monk",
		Level: 90,
		Race:  "Lalafell",
		Date:  time.Now(),
	}
	party["Siryu"] = Character{
		Name:  "Siryu Unslay",
		Job:   "Dark Knight",
		Level: 90,
		Race:  "Au'Ra",
		Date:  time.Now(),
	}
	party["James"] = Character{
		Name:  "James Carmike",
		Job:   "Sage",
		Level: 90,
		Race:  "Viera",
		Date:  time.Now(),
	}
	party["Kisa"] = Character{
		Name:  "Kisa Rozen",
		Job:   "Reaper",
		Level: 90,
		Race:  "Hyur",
		Date:  time.Now(),
	}
	// for _, member := range party {
	// 	fmt.Println(member)
	// }

	output, err := json.Marshal(party)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var lightParty map[string]Character

	err = json.Unmarshal(output, &lightParty)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(lightParty["Siryu"])

}
