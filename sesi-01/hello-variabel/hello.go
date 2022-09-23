package main

import "fmt"

func main() {

	//belajar variabel with data type
	var nama string = "Dinda"
	var umur int = 21
	fmt.Println("Ini adalah namanya ==>", nama)
	fmt.Println("Ini adalah umurnya ==>", umur)

	//belajar variabel without data type
	var namaAnda = "Mikasa"
	var umurAnda = 22

	fmt.Printf("%T, %T", namaAnda, umurAnda)

	//belajar short declaration
	namaSaya := "Dinda"
	umurSaya := 21

	fmt.Printf("%T, %T", namaSaya, umurSaya)

	//belajar multiple variable declarations
	var eren, mikasa, armin string = "1", "2", "3"
	fmt.Println("test data multiple", eren, mikasa, armin)

	//belajar multiple variable short declarations
	var jean, connie, sasha = "survey corps", "military police", "garrison"
	levi, erwin, hange := "shiganshina", 235, "commander"
	fmt.Println(jean, connie, sasha)
	fmt.Println(levi, erwin, hange)

	//belajar underscore variable
	var firstVariable string
	var name, age, address = "Dinda", 21, "Jalan merpati"
	_, _, _, _ = firstVariable, name, age, address
	fmt.Printf("test underscore variable > %T\n", name)
	fmt.Printf("halo nama saya %s, umur saya yaitu %d, dan saya tinggal di %s\n", name, age, address)

	//belajar fmt.Printf usage
	first, second := 1, "2"

	fmt.Printf("Tipe data variabel first adalah %T \n", first)
	fmt.Printf("Tipe data variabel second adalah %T \n", second)

	//belajar fmt.Printf usage
	var namaDia = "Ecen"

	var umurDia = 12

	var alamatDia = "Jalan BNI City"

	fmt.Printf("perkenalkan namanya %s, umurnya adalah %d, dan ia tinggal di %s", namaDia, umurDia, alamatDia)

}
