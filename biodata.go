// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// type Friend struct {
// 	Nama string
// 	Alamat string
// 	Pekerjaan string
// 	Alasan string
// }

// func main() {
// 	friends := []Friend {
// 		{Nama: "Teman 1", Alamat: "Jakarta", Pekerjaan: "kerja", Alasan: "BU"},
// 		{Nama: "Teman 2", Alamat: "Jakarta", Pekerjaan: "kerja", Alasan: "BU"},
// 		{Nama: "Teman 3", Alamat: "Jakarta", Pekerjaan: "kerja", Alasan: "BU"},
// 	}
// 	args := os.Args[1:]
// 	fmt.Println(args)
// 	numbr,_ := strconv.Atoi(args[0])

// 	var printByArg = func(i int) {
// 		fmt.Println(friends[i - 1])
// 	}

// 	printByArg(numbr)
// }