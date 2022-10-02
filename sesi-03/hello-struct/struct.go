package main

import "fmt"

type Employee struct {
	nama   string
	umur   int
	divisi string
}

type Employee1 struct {
	name     string
	age      int
	division string
}

func main() {
	//struct (giving value to struct)
	var employee Employee

	employee.nama = "Levi Ackerman"

	employee.umur = 28

	employee.divisi = "Captain Squad Survey Corps"

	fmt.Println(employee.nama)
	fmt.Println(employee.umur)
	fmt.Println(employee.divisi)
}
