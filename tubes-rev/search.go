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
	fmt.Println("| ")
}

func sequentialSearch(A tabTest, n, x int, find, userID string) {
	var found bool = false
	var j int = 1
	switch n {
	case 1:
		for i := 0; i <= x; i++ {
			if find == A[i].hasilD && userID == A[i].userID {
				fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
				found = true
				j++
			}
		}
	case 2:
		for i := 0; i <= x; i++ {
			if find == A[i].hasilA && userID == A[i].userID {
				fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
				found = true
				j++
			}
		}
	case 3:
		for i := 0; i <= x; i++ {
			if find == A[i].hasilS && userID == A[i].userID {
				fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
				found = true
				j++
			}
		}
	}
	if !found {
		fmt.Println("|         Data tidak ditemukan        |")
	}
}

func binarySearch(A tabTest, x, date, year int, month, userID string) {
	var med, kr, kn int = 0, 0, x - 1
	var j int = 1
	var found bool
	for i := 0; i <= x; i++ {
		found = false
		for kr <= kn && !found {
			med = (kr + kn) / 2
			if date > A[med].date {
				kn = med - 1
			} else if date < A[med].date {
				kr = med + 1
			} else if A[med].month == month && A[med].year == year && userID == A[i].userID {
				fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
				found = true
				j++
			}
		}
	}
	if !found {
		fmt.Println("|         Data tidak ditemukan        |")
	}
}
