package main

import (
	"fmt"
)

func loginPage(A *tabTest, B arrPengguna, maxU, m int, userID string) {
	//I.S.
	//F.S.
	var x, idx int
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
		test(A, B, x, idx, maxU, &m, userID)
	case 2:
		riwayat(A, B, maxU, x, m, &idx, userID)
	case 3:
		changeData(A, B, maxU, x, m, idx, userID)
	case 4:
		deleteData(A, B, maxU, x, idx, &m, userID)
	case 5:
		laporan(A, B, maxU, m, userID)
	case 6:
		ui(A, &B, &maxU, m)
	default:
	}
}
