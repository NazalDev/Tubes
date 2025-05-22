package main

import (
	"fmt"
)

func riwayat(A arrPengguna, temp, userIdx int, idx *int) {
	var a, b, pick, d, y int
	var find string = "null"
	if A[userIdx].totTest < 0 {
		fmt.Println("=======================================")
		fmt.Println("           Data Not Yet Made           ")
		fmt.Println("=======================================")
		fmt.Println("|                                     |")
	} else {
		fmt.Println("| Riwayat dan List Data :")
		for i := 0; i < A[userIdx].totTest; i++ {
			fmt.Printf("| %d. %d %s %d %s %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].userID, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
		}
		if temp == 2 {
			fmt.Println("| ")
			fmt.Println("| Urutkan atau cari data ?  ")
			fmt.Println("| 1. Urutkan")
			fmt.Println("| 2. Cari data")
			fmt.Println("| 3. Kembali")
			fmt.Print("| Pilih angka (1/2/3) ")
			fmt.Scan(&pick)
			fmt.Println("| ")
			if pick == 1 {
				sortMenu(&a, &b)
				switch b {
				case 1:
					if a == 4 {
						sortDateAsc(A, userIdx, d, y, find)
					} else {
						sortSkorAsc(A, userIdx, a, find)
					}
				case 2:
					if a == 4 {
						sortDateDes(A, userIdx)
					} else {
						sortSkorDes(A, userIdx, a)
					}
				}
			} else if pick == 2 {
				searchMenu(&b)
				switch b {
				case 1:
					fmt.Print("| Tanggal : ")
					fmt.Scan(&d)
					fmt.Print("| Bulan (Depan wajib kapital) : ")
					fmt.Scan(&find)
					fmt.Print("| Tahun : ")
					fmt.Scan(&y)
					sortDateAsc(A, userIdx, d, y, find)
				case 2:
					fmt.Println("| Data yang akan di search :")
					fmt.Println("| 1. Depresi")
					fmt.Println("| 2. Cemas")
					fmt.Println("| 3. Stress")
					fmt.Print("| Pilih angka (1/2/3) ")
					fmt.Scan(&a)
					fmt.Println("|")
					fmt.Println("| Level data yang akan dicari (Normal/Ringan/Sedang/Berat/Sangat_Berat)")
					fmt.Println("| Note : Sesuaikan penulisan dengan contoh")
					fmt.Print("| Pilih level : ")
					fmt.Scan(&find)
					fmt.Println("|")
					sortSkorAsc(A, userIdx, a, find)
				}
			}
		}
	}
	if temp == 3 || temp == 4 {
		fmt.Print("| Pilih data: ")
		fmt.Scan(idx)
		fmt.Println()
	}
}
