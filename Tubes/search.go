package main

import (
	"fmt"
)

func searchMenu(pick *int) {
	fmt.Println("=========== Pencarian Data  ===========")
	fmt.Println("| 1. Berdasarkan tanggal")
	fmt.Println("| 2. Berdasarkan skor")
	fmt.Print("| Pilih angka (1/2) ")
	fmt.Scan(pick)
	fmt.Println("|")
}

func sequentialSearch(A arrPengguna, userIdx, n int, find string) {
	var found bool = false
	var j int = 1
	switch n {
	case 1:
		for i := 0; i < A[userIdx].totTest; i++ {
			if find == A[userIdx].tabTest[i].hasilD {
				fmt.Printf("| %d. %d %s %d %s %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", j, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].userID, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
				found = true
				j++
			}
		}
	case 2:
		for i := 0; i < A[userIdx].totTest; i++ {
			if find == A[userIdx].tabTest[i].hasilA {
				fmt.Printf("| %d. %d %s %d %s %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", j, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].userID, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
				found = true
				j++
			}
		}
	case 3:
		for i := 0; i < A[userIdx].totTest; i++ {
			if find == A[userIdx].tabTest[i].hasilS {
				fmt.Printf("| %d. %d %s %d %s %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", j, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].userID, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
				found = true
				j++
			}
		}
	}
	if !found {
		fmt.Println("|         Data tidak ditemukan        |")
	}
}

func dateSearch(A arrPengguna, userIdx, date, year int, month string) {
	var found bool = false
	var j int = 1
	for i := 0; i < A[userIdx].totTest; i++ {
		if date == A[userIdx].tabTest[i].date {
			if month == A[userIdx].tabTest[i].month && year == A[userIdx].tabTest[i].year {
				fmt.Printf("| %d. %d %s %d %s %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", j, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].userID, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
				found = true
				j++
			}
		}
	}
	if !found {
		fmt.Println("|         Data tidak ditemukan        |")
	}
}

func binarySearch(A arrPengguna, n int, find string) int {
	var found int = -1
	var med, kr, kn int = 0, 0, n - 1
	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if A[med].userId == find {
			found = med
		} else if A[med].userId < find {
			kr = med + 1
		} else {
			kn = med - 1
		}
	}
	return found
}
