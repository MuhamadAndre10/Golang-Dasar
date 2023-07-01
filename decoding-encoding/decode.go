package decodingencoding

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	City       string `json:"city"`
	Occupation string `json:"occupation"`
}

func Decode() {
	jsonString := `{"name":"John","age":30,"city":"New York"}`

	var person Person
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	switch person.Occupation {
	case "":
		fmt.Println("nil")
	default:
		fmt.Println("error")
	}

	fmt.Println(person.Name)
	fmt.Println(person.Age)
	fmt.Println(person.City)
	fmt.Printf("%T", person.Occupation)
}
