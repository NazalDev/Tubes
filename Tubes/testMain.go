package main

import (
	"fmt"
)

func loginPage(A *arrPengguna, maxU, userIdx int) {
	//I.S.
	//F.S.
	var x, idx int
	for {
		fmt.Println("=============   Menu Test   ===========")
		fmt.Println("|1. Start Test                        |")
		fmt.Println("|2. Riwayat                           |")
		fmt.Println("|3. Ubah Data                         |")
		fmt.Println("|4. Hapus Data                        |")
		fmt.Println("|5. Laporan                           |")
		fmt.Println("|6. Back                              |")
		fmt.Println("|7. Exit                              |")
		fmt.Println("=======================================")
		fmt.Println()
		fmt.Print("Pilih salah satu menu : ")
		fmt.Scan(&x)
		switch x {
		case 1:
			test(&A[userIdx].tabTest[A[userIdx].totTest])
			A[userIdx].totTest += 1
			fmt.Println(A[userIdx].totTest, userIdx)
		case 2:
			riwayat(*A, x, userIdx, &idx)
		case 3:
			changeData(A, userIdx, x, idx)
		case 4:
			deleteData(A, userIdx, x, idx)
		case 5:
			laporan(*A, userIdx)
		case 6:
			ui(A, &maxU)
		default:
		}
		if x == 7 || x == 6 {
			break
		}
	}
}
