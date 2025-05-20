package main

import (
	"fmt"
)

func sortSkorAsc(A tabTest, n, x int, find, userID string) {
	var temp dataTes
	var i, idxMin int
	var j int = 1
	switch n {
	case 1:
		for i = 0; i < x; i++ {
			idxMin = i
			for j := i + 1; j <= x; j++ {
				if A[idxMin].depress > A[j].depress {
					idxMin = j
				}
			}
			temp = A[idxMin]
			A[idxMin] = A[i]
			A[i] = temp
		}
	case 2:
		for i = 0; i < x; i++ {
			idxMin = i
			for j := i + 1; j <= x; j++ {
				if A[idxMin].anxiety > A[j].anxiety {
					idxMin = j
				}
			}
			temp = A[idxMin]
			A[idxMin] = A[i]
			A[i] = temp
		}
	case 3:
		for i = 0; i < x; i++ {
			idxMin = i
			for j := i + 1; j <= x; j++ {
				if A[idxMin].stress > A[j].stress {
					idxMin = j
				}
			}
			temp = A[idxMin]
			A[idxMin] = A[i]
			A[i] = temp
		}
	}
	if find == "null" {
		fmt.Println("| Hasil Mengurutkan :")
		for i = 0; i <= x; i++ {
			if userID == A[i].userID {
				fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", j, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
				j++
			}
		}
	} else {
		sequentialSearch(A, n, x, find, userID)
	}
}
func sortSkorDes(A tabTest, n, x int, userID string) {
	var temp dataTes
	var i, idxMin int
	var j int = 1
	switch n {
	case 1:
		for i = 0; i < x; i++ {
			idxMin = i
			for j := i + 1; j <= x; j++ {
				if A[idxMin].depress < A[j].depress {
					idxMin = j
				}
			}
			temp = A[idxMin]
			A[idxMin] = A[i]
			A[i] = temp
		}
	case 2:
		for i = 0; i < x; i++ {
			idxMin = i
			for j := i + 1; j <= x; j++ {
				if A[idxMin].anxiety < A[j].anxiety {
					idxMin = j
				}
			}
			temp = A[idxMin]
			A[idxMin] = A[i]
			A[i] = temp
		}
	case 3:
		for i = 0; i < x; i++ {
			idxMin = i
			for j := i + 1; j <= x; j++ {
				if A[idxMin].stress < A[j].stress {
					idxMin = j
				}
			}
			temp = A[idxMin]
			A[idxMin] = A[i]
			A[i] = temp
		}
	}
	fmt.Println("| Hasil Mengurutkan :")
	for i = 0; i <= x; i++ {
		if userID == A[i].userID {
			fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", j, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
			j++
		}
	}
}

func sortDateAsc(A tabTest, x, date, year int, month, userID string) {
	var i, j int
	var temp dataTes
	i = 1
	for i <= x {
		j = i
		temp = A[j]
		for j > 0 && temp.date < A[j-1].date {
			A[j] = A[j-1]
			j -= 1
		}
		A[j] = temp
		i += 1
	}
	if month == "null" {
		fmt.Println("| Hasil Mengurutkan :")
		for i = 0; i <= x; i++ {
			if userID == A[i].userID {
				fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", j, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
				j++
			}
		}
	} else {
		binarySearch(A, x, date, year, month, userID)
	}
}

func sortDateDes(A tabTest, x int, userID string) {
	var i, j int
	var temp dataTes
	i = 1
	for i <= x {
		j = i
		temp = A[j]
		for j > 0 && temp.date > A[j-1].date {
			A[j] = A[j-1]
			j -= 1
		}
		A[j] = temp
		i += 1
	}
	fmt.Println("| Hasil Mengurutkan :")
	j = 0
	for i = 0; i <= x; i++ {
		if userID == A[i].userID {
			fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", j, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
			j++
		}
	}
}

func sortMenu(a, b *int) {
	fmt.Println("| Urutkan sesuai :")
	fmt.Println("| 1. Level Depresi")
	fmt.Println("| 2. Level Kecemasan")
	fmt.Println("| 3. Level Stress")
	fmt.Println("| 4. Tanggal pengerjaan")
	fmt.Print("| Pilih angka (1/2/3/4)  ")
	fmt.Scan(a)
	fmt.Println("| ")
	fmt.Println("| Menaik/menurun? ")
	fmt.Println("| 1. Menaik")
	fmt.Println("| 2. Menurun")
	fmt.Print("| Pilih angka (1/2)  ")
	fmt.Scan(b)
	fmt.Println()
}
