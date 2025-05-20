package main

import (
	"fmt"
)

func riwayat(A *tabTest, B arrPengguna, maxU, temp, x int, idx *int, userID string) {
	var find string = "null"
	var a, b, pick, d, y int
	var j, i int = 1, 0
	if x < 0 {
		fmt.Println("=======================================")
		fmt.Println("           Data Not Yet Made           ")
		fmt.Println("=======================================")
		fmt.Println("|                                     |")
		loginPage(A, B, maxU, x, userID)
	} else {
		fmt.Println("| Riwayat dan List Data :")
		for i := 0; i <= x; i++ {
			if userID == A[i].userID {
				fmt.Printf("| %d. %d %s %d %s %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", j, A[i].date, A[i].month, A[i].year, A[i].time, A[i].userID, A[i].hasilD, A[i].hasilA, A[i].hasilS)
				j++
			}
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
						sortDateAsc(*A, x, d, y, find, userID)
					} else {
						sortSkorAsc(*A, a, x, find, userID)
					}
				case 2:
					if a == 4 {
						sortDateDes(*A, x, userID)
					} else {
						sortSkorDes(*A, a, x, userID)
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
					sortDateAsc(*A, x, d, y, find, userID)
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
					sortSkorAsc(*A, a, x, find, userID)
				}
			}
			loginPage(A, B, maxU, x, userID)
		}
	}
	if temp == 3 || temp == 4 {
		fmt.Print("| Pilih data: ")
		fmt.Scan(idx)
		fmt.Println()
		for {
			if userID == A[i].userID {
				j++
			}
			if j == *idx || i > x {
				idx = &i
				break
			} else {
				i++
			}
		}
	}
}
