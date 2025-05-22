package main

import (
	"fmt"
)

func sortSkorAsc(A arrPengguna, userIdx, n int, find string) {
	var temp dataTes
	var i, idxMin int
	switch n {
	case 1:
		for i = 0; i < A[userIdx].totTest-1; i++ {
			idxMin = i
			for j := i + 1; j < A[userIdx].totTest; j++ {
				if A[userIdx].tabTest[idxMin].depress > A[userIdx].tabTest[j].depress {
					idxMin = j
				}
			}
			temp = A[userIdx].tabTest[idxMin]
			A[userIdx].tabTest[idxMin] = A[userIdx].tabTest[i]
			A[userIdx].tabTest[i] = temp
		}
	case 2:
		for i = 0; i < A[userIdx].totTest-1; i++ {
			idxMin = i
			for j := i + 1; j < A[userIdx].totTest; j++ {
				if A[userIdx].tabTest[idxMin].anxiety > A[userIdx].tabTest[j].anxiety {
					idxMin = j
				}
			}
			temp = A[userIdx].tabTest[idxMin]
			A[userIdx].tabTest[idxMin] = A[userIdx].tabTest[i]
			A[userIdx].tabTest[i] = temp
		}
	case 3:
		for i = 0; i < A[userIdx].totTest-1; i++ {
			idxMin = i
			for j := i + 1; j < A[userIdx].totTest; j++ {
				if A[userIdx].tabTest[idxMin].stress > A[userIdx].tabTest[j].stress {
					idxMin = j
				}
			}
			temp = A[userIdx].tabTest[idxMin]
			A[userIdx].tabTest[idxMin] = A[userIdx].tabTest[i]
			A[userIdx].tabTest[i] = temp
		}
	}
	if find == "null" {
		fmt.Println("| Hasil Mengurutkan :")
		for i = 0; i < A[userIdx].totTest; i++ {
			fmt.Printf("| %d. %d %s %d %s %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].userID, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
		}
	} else {
		sequentialSearch(A, userIdx, n, find)
	}
}
func sortSkorDes(A arrPengguna, userIdx, n int) {
	var temp dataTes
	var i, idxMin int
	switch n {
	case 1:
		for i = 0; i < A[userIdx].totTest-1; i++ {
			idxMin = i
			for j := i + 1; j < A[userIdx].totTest; j++ {
				if A[userIdx].tabTest[idxMin].depress < A[userIdx].tabTest[j].depress {
					idxMin = j
				}
			}
			temp = A[userIdx].tabTest[idxMin]
			A[userIdx].tabTest[idxMin] = A[userIdx].tabTest[i]
			A[userIdx].tabTest[i] = temp
		}
	case 2:
		for i = 0; i < A[userIdx].totTest-1; i++ {
			idxMin = i
			for j := i + 1; j < A[userIdx].totTest; j++ {
				if A[userIdx].tabTest[idxMin].anxiety < A[userIdx].tabTest[j].anxiety {
					idxMin = j
				}
			}
			temp = A[userIdx].tabTest[idxMin]
			A[userIdx].tabTest[idxMin] = A[userIdx].tabTest[i]
			A[userIdx].tabTest[i] = temp
		}
	case 3:
		for i = 0; i < A[userIdx].totTest-1; i++ {
			idxMin = i
			for j := i + 1; j < A[userIdx].totTest; j++ {
				if A[userIdx].tabTest[idxMin].stress < A[userIdx].tabTest[j].stress {
					idxMin = j
				}
			}
			temp = A[userIdx].tabTest[idxMin]
			A[userIdx].tabTest[idxMin] = A[userIdx].tabTest[i]
			A[userIdx].tabTest[i] = temp
		}
	}
	fmt.Println("| Hasil Mengurutkan :")
	for i = 0; i < A[userIdx].totTest; i++ {
		fmt.Printf("| %d. %d %s %d %s %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].userID, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
	}
}

func sortDateAsc(A arrPengguna, userIdx, date, year int, month string) {
	var i, j int
	var temp dataTes
	i = 1
	for i < A[userIdx].totTest {
		j = i
		temp = A[userIdx].tabTest[j]
		for j > 0 && temp.date < A[userIdx].tabTest[j-1].date {
			A[userIdx].tabTest[j] = A[userIdx].tabTest[j-1]
			j -= 1
		}
		A[userIdx].tabTest[j] = temp
		i += 1
	}
	if month == "null" {
		fmt.Println("| Hasil Mengurutkan :")
		for i = 0; i < A[userIdx].totTest; i++ {
			fmt.Printf("| %d. %d %s %d %s %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].userID, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
		}
	} else {
		dateSearch(A, userIdx, date, year, month)
	}
}

func sortDateDes(A arrPengguna, userIdx int) {
	var i, j int
	var temp dataTes
	i = 1
	for i < A[userIdx].totTest {
		j = i
		temp = A[userIdx].tabTest[j]
		for j > 0 && temp.date > A[userIdx].tabTest[j-1].date {
			A[userIdx].tabTest[j] = A[userIdx].tabTest[j-1]
			j -= 1
		}
		A[userIdx].tabTest[j] = temp
		i += 1
	}
	fmt.Println("| Hasil Mengurutkan :")
	for i = 0; i < A[userIdx].totTest; i++ {
		fmt.Printf("| %d. %d %s %d %s %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].userID, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
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
