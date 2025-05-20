package main

import "fmt"

func changeData(A *tabTest, B arrPengguna, maxU, temp, x, idx int, userID string) {
	var a string
	if idx <= x && idx >= 0 {
		riwayat(A, B, maxU, temp, x, &idx, userID)
		confirm(&a)
		if a == "y" || a == "Y" {
			test(A, B, maxU, temp, idx, &x, userID)
		}
	}
}

func deleteData(A *tabTest, B arrPengguna, maxU, temp, idx int, x *int, userID string) {
	var a string
	riwayat(A, B, maxU, temp, *x, &idx, userID)
	confirm(&a)
	if a == "y" || a == "Y" {
		for i := idx - 1; i <= *x; i++ {
			A[i] = A[i+1]
			A[i].testID--
		}
		*x -= 1
		riwayat(A, B, maxU, 2, *x, &idx, userID)
	}
}

func confirm(answer *string) {
	fmt.Println("=======================================")
	fmt.Println("|            Yakin ? Y / N            |")
	fmt.Println("=======================================")
	fmt.Scan(answer)
}
