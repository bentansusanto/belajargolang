package main

import "fmt"

func main(){
	tampil_mahasiswa(182,178)
}

func tampil_mahasiswa(x int, y int){
	var Aldo = map[string]int{"tinggi":x};
	var Yosep = map[string]int{"tinggi":y};
	fmt.Println("Aldo :", Aldo["tinggi"],"cm")
	fmt.Println("Yosep :", Yosep["tinggi"],"cm")
}
