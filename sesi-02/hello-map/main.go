package main

import (
	"fmt"
)

func main() {
	//map
	var person map[string]string // Deklarasi

	person = map[string]string{} // Inisialisasi

	person["name"] = "Dinda"

	person["age"] = "21"

	person["address"] = "Jalan Merpati"

	fmt.Println("name:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"])

	//map (looping with map)
	var dataPerson = map[string]string{
		"dataName":    "Aiden",
		"dataAge":     "17",
		"dataAddress": "Senopati",
	}

	for key, value := range dataPerson {
		fmt.Println(key, ":", value)
	}

	//map (deleting value)
	var person1 = map[string]string{
		"name1":    "Raysie",
		"age1":     "19",
		"Address1": "Depok",
	}

	fmt.Println("Before deleting:", person1)

	delete(person1, "age1")

	fmt.Println("After deleting:", person1)

	//map (detecting a value)
	var person2 = map[string]string{
		"name2":    "Sky",
		"age2":     "20",
		"Address2": "Bekasi Timur",
	}

	value, exist := person2["age2"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value doesn't exist")
	}

	delete(person2, "age2")

	value, exist = person2["age2"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}

	//map (combining slice with map)
	var people = []map[string]string{
		{"peopleName": "Dinda", "peopleAge": "23"},
		{"peopleName": "Farah", "peopleAge": "24"},
		{"peopleName": "Delvia", "peopleAge": "25"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, peopleName: %s, peopleAge: %s\n", i, person["peopleName"], person["peopleAge"])
	}
}
