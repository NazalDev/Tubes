package main

import "fmt"

func changeData(A *arrPengguna, userIdx, temp, idx int) {
	var a string
	if idx < A[userIdx].totTest && idx >= 0 {
		riwayat(*A, temp, userIdx, &idx)
		idx--
		confirm(&a)
		if a == "y" || a == "Y" {
			test(&A[userIdx].tabTest[idx])
		}
	}
}

func deleteData(A *arrPengguna, userIdx, temp, idx int) {
	var a string
	riwayat(*A, temp, userIdx, &idx)
	confirm(&a)
	if a == "y" || a == "Y" {
		for i := idx - 1; i < A[userIdx].totTest; i++ {
			A[userIdx].tabTest[i] = A[userIdx].tabTest[i+1]
		}
		A[userIdx].totTest--
		riwayat(*A, 2, userIdx, &idx)
	}
}

func confirm(answer *string) {
	fmt.Println("=======================================")
	fmt.Println("|            Yakin ? Y / N            |")
	fmt.Println("=======================================")
	fmt.Scan(answer)
}
