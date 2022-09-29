package main

import (
	"fmt"
	"os"
	"strconv"
)

type Peserta struct {
	noAbsensi int
	nama      string
	alamat    string
	umur      int
	divisi    string
	alasan    string
}

func main() {
	//data peserta
	var peserta = []Peserta{
		{noAbsensi: 1, nama: "Dinda", alamat: "Jalan Merpati", umur: 21, divisi: "Teacher Assistant", alasan: "Ingin belajar IT"},
		{noAbsensi: 2, nama: "Ayu", alamat: "Jalan Halim P.K", umur: 20, divisi: "Curriculum Development", alasan: "Ingin mengetahui bahasa Go"},
		{noAbsensi: 3, nama: "Hikmah", alamat: "Jurumudi", umur: 22, divisi: "Digital Marketing", alasan: "Menambah khasanah"},
		{noAbsensi: 4, nama: "Pertiwi", alamat: "Benda", umur: 23, divisi: "SEO Specialist", alasan: "Memperluas pengetahuan pemrograman"},
	}

	//argument
	var argsRaw = os.Args
	var args = argsRaw[1]
	noAbsensi, err := strconv.Atoi(args)
	_ = err

	getPeserta(peserta, noAbsensi)
}

func getPeserta(p []Peserta, noAbsensi int) {

	if noAbsensi <= len(p) {
		fmt.Println("Data Peserta:")
		fmt.Println("noAbsensi:", p[noAbsensi-1].noAbsensi)
		fmt.Println("nama:", p[noAbsensi-1].nama)
		fmt.Println("alamat:", p[noAbsensi-1].alamat)
		fmt.Println("umur:", p[noAbsensi-1].umur)
		fmt.Println("divisi:", p[noAbsensi-1].divisi)
		fmt.Println("alasan:", p[noAbsensi-1].alasan)
	} else {
		fmt.Println("Data peserta tertanda")
	}
}
