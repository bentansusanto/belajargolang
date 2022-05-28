package main

import "fmt"

// method
func main(){
	var s1 pelajar
	var s2 = pelajar{"yosep",18}
	s1.nama = "aldo"
	s1.umur = 19
	s1.namasaya()
	s2.namasaya()
}

type pelajar struct {   
	//nama variabel dan tipe data
	nama string
	umur int
}

func (s pelajar) namasaya(){
	// menampilkan data dari method main ke dalam method namasaya
	fmt.Println("nama saya adalaah",s.nama)
	fmt.Println("umur saya adalaah",s.umur)
}


// struct
/*func main(){
	var x1 pelajar
	x1.nama = "johan"
	x1.kelas = 12 
	x1.umur = 17
	fmt.Println("nama :", x1.nama)
	fmt.Println("kelas :", x1.kelas)
	fmt.Println("umur :", x1.umur)
}

type pelajar struct {   
	//nama variabel dan tipe data
	nama string
	kelas int 
	umur int 
}


// pointer
/* func main(){
	var nomor int = 10;
	var alamat_nomor *int = &nomor
	var tas string = "spiderman"
	var alamat_tas *string = &tas

	nomor = 30;
	fmt.Println("Nilai dari nomor :", nomor)
	fmt.Println("Alamat Variabel Nomor :", alamat_nomor)
	fmt.Println("Nilai dari tas :", tas)
	fmt.Println("Alamat Variabel Tas L :", alamat_tas)
*/
